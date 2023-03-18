package dev.emortal.api.utils.kafka;

import com.google.common.util.concurrent.ThreadFactoryBuilder;
import com.google.protobuf.AbstractMessage;
import com.google.protobuf.InvalidProtocolBufferException;
import dev.emortal.api.utils.parser.MessageProtoConfig;
import dev.emortal.api.utils.parser.MessagingService;
import dev.emortal.api.utils.parser.ProtoParserRegistry;
import org.apache.kafka.clients.consumer.ConsumerConfig;
import org.apache.kafka.clients.consumer.ConsumerRecord;
import org.apache.kafka.clients.consumer.ConsumerRecords;
import org.apache.kafka.clients.consumer.KafkaConsumer;
import org.apache.kafka.common.header.Header;
import org.apache.kafka.common.header.Headers;
import org.jetbrains.annotations.NotNull;
import org.jetbrains.annotations.Nullable;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.time.Duration;
import java.util.Collections;
import java.util.HashSet;
import java.util.Map;
import java.util.Properties;
import java.util.Set;
import java.util.concurrent.ConcurrentHashMap;
import java.util.function.Consumer;

public class KafkaCore {
    private static final Logger LOGGER = LoggerFactory.getLogger(KafkaCore.class);

    private final @NotNull KafkaConsumer<String, byte[]> consumer;

    private final Map<Class<?>, Consumer<AbstractMessage>> protoListeners = new ConcurrentHashMap<>();
    private final @NotNull Set<String> consumedTopics = Collections.synchronizedSet(new HashSet<>());

    private final Thread consumerThread;

    public KafkaCore(@NotNull KafkaSettings settings) {
        Properties properties = new Properties();

        properties.put(ConsumerConfig.BOOTSTRAP_SERVERS_CONFIG, settings.getBootstrapServers());
        properties.put(ConsumerConfig.CLIENT_ID_CONFIG, settings.getClientId());
        properties.put(ConsumerConfig.GROUP_ID_CONFIG, settings.getGroupId());
        properties.put(ConsumerConfig.ENABLE_AUTO_COMMIT_CONFIG, settings.isAutoCommit());

        this.consumer = new KafkaConsumer<>(properties);

        new ThreadFactoryBuilder().setNameFormat("kafka-consumer").build();
        this.consumerThread = new Thread(this::consume, "kafka-consumer");
    }

    private void consume() {
        while (!Thread.currentThread().isInterrupted()) {
            ConsumerRecords<String, byte[]> records = this.consumer.poll(Duration.ofSeconds(5));

            for (ConsumerRecord<String, byte[]> record : records) {
                String protoType = this.getProtoType(record.headers());
                if (protoType == null) {
                    LOGGER.warn("Received message without X-Proto-Type header");
                    continue;
                }

                AbstractMessage message;
                try {
                    message = ProtoParserRegistry.parse(protoType, record.value());
                } catch (InvalidProtocolBufferException e) {
                    LOGGER.warn("Failed to parse message", e);
                    continue;
                }

                try {
                    this.protoListeners.get(message.getClass()).accept(message);
                } catch (Exception e) {
                    LOGGER.warn("Failed to handle message", e);
                }
            }
        }
    }

    private @Nullable String getProtoType(@NotNull Headers headers) {
        for (Header header : headers) {
            if (header.key().equals("X-Proto-Type")) {
                return new String(header.value());
            }
        }
        return null;
    }

    @SuppressWarnings("unchecked")
    public <T extends AbstractMessage> void setListener(@NotNull Class<T> messageType, @NotNull Consumer<T> listener) {
        MessageProtoConfig<T> protoConfig = ProtoParserRegistry.getParser(messageType);
        if (protoConfig == null) {
            throw new IllegalArgumentException("No parser found for " + messageType.getName());
        }

        if (protoConfig.service() != MessagingService.KAFKA) {
            throw new IllegalArgumentException("Parser for " + messageType.getName() + " is not for Kafka");
        }

        if (this.protoListeners.containsKey(messageType)) {
            throw new IllegalArgumentException("Listener already set for " + messageType.getName());
        }

        if (!this.consumedTopics.contains(protoConfig.topic())) {
            LOGGER.debug("Subscribing to topic {}", protoConfig.topic());
            this.consumedTopics.add(protoConfig.topic());
            this.consumer.subscribe(this.consumedTopics);
        }

        this.protoListeners.put(messageType, (Consumer<AbstractMessage>) listener);
    }

    public void shutdown() {
        this.consumer.close();
        this.consumerThread.interrupt();
    }
}
