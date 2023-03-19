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
import org.apache.kafka.common.serialization.ByteArrayDeserializer;
import org.apache.kafka.common.serialization.StringDeserializer;
import org.jetbrains.annotations.NotNull;
import org.jetbrains.annotations.Nullable;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.time.Duration;
import java.time.Instant;
import java.util.Collections;
import java.util.HashSet;
import java.util.Map;
import java.util.Properties;
import java.util.Set;
import java.util.UUID;
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

        // Nullable options
        if (settings.getAutoCommit() != null) properties.put(ConsumerConfig.ENABLE_AUTO_COMMIT_CONFIG, settings.getAutoCommit());

        String groupId = settings.getGroupId();
        if (groupId == null) groupId = "kafka-core-" + UUID.randomUUID();
        properties.put(ConsumerConfig.GROUP_ID_CONFIG, groupId);

        System.out.println("properties: " + properties);

        this.consumer = new KafkaConsumer<>(properties, new StringDeserializer(), new ByteArrayDeserializer());

        new ThreadFactoryBuilder().setNameFormat("kafka-consumer").build();
        this.consumerThread = new Thread(this::consume, "kafka-consumer");
    }

    private void consume() {
        LOGGER.info("Starting Kafka consumer thread");

        Instant lastPoll = Instant.now().minusSeconds(5);
        while (!Thread.currentThread().isInterrupted()) {
            // Only poll every second
            Duration timeSinceLastPoll = Duration.between(lastPoll, Instant.now());
            if (timeSinceLastPoll.toMillis() < 1000) {
                try {
                    Thread.sleep(1000 - timeSinceLastPoll.toMillis());
                } catch (InterruptedException e) {
                    LOGGER.warn("Interrupted while sleeping", e);
                }
            }

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

        if (!this.consumerThread.isAlive()) {
            this.consumerThread.start();
        }

        this.protoListeners.put(messageType, (Consumer<AbstractMessage>) listener);
    }

    public void shutdown() {
        this.consumer.close();
        this.consumerThread.interrupt();
    }
}
