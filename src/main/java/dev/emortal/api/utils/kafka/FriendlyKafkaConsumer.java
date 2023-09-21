package dev.emortal.api.utils.kafka;

import com.google.common.util.concurrent.ThreadFactoryBuilder;
import com.google.protobuf.AbstractMessage;
import com.google.protobuf.InvalidProtocolBufferException;
import dev.emortal.api.utils.parser.MessageProtoConfig;
import dev.emortal.api.utils.parser.ProtoParserRegistry;
import org.apache.kafka.clients.consumer.ConsumerRecord;
import org.apache.kafka.clients.consumer.ConsumerRecords;
import org.apache.kafka.clients.consumer.KafkaConsumer;
import org.apache.kafka.common.TopicPartition;
import org.apache.kafka.common.header.Header;
import org.apache.kafka.common.header.Headers;
import org.apache.kafka.common.serialization.ByteArrayDeserializer;
import org.apache.kafka.common.serialization.StringDeserializer;
import org.jetbrains.annotations.NotNull;
import org.jetbrains.annotations.Nullable;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.time.Duration;
import java.util.Collections;
import java.util.HashSet;
import java.util.Map;
import java.util.Set;
import java.util.concurrent.ConcurrentHashMap;
import java.util.concurrent.Executors;
import java.util.concurrent.ScheduledExecutorService;
import java.util.concurrent.TimeUnit;
import java.util.concurrent.atomic.AtomicInteger;
import java.util.function.Consumer;

public final class FriendlyKafkaConsumer {
    private static final Logger LOGGER = LoggerFactory.getLogger(FriendlyKafkaConsumer.class);
    private static final int POLL_TIMEOUT_MS = 250;
    private static final Duration POLL_TIMEOUT = Duration.ofMillis(POLL_TIMEOUT_MS);

    private static final AtomicInteger GROUP_COUNTER = new AtomicInteger(0);

    private final @NotNull KafkaConsumer<String, byte[]> consumer;
    private final boolean hasConsumerGroup;

    private final Map<Class<?>, Set<Consumer<AbstractMessage>>> protoListeners = new ConcurrentHashMap<>();
    private final Set<String> consumedTopics = new HashSet<>();
    // Only used if there isn't a consumer group - so we manually set the partitions (not subscribe)
    private final Set<TopicPartition> partitions = Collections.synchronizedSet(new HashSet<>());

    private final ScheduledExecutorService scheduler = Executors.newSingleThreadScheduledExecutor(
            new ThreadFactoryBuilder().setNameFormat("kafka-consumer-scheduler")
                    .setUncaughtExceptionHandler((thread, exception) -> LOGGER.error("Error in Kafka consumer: ", exception))
                    .build());

    private final boolean autoCommit;

    public FriendlyKafkaConsumer(@NotNull KafkaSettings settings) {
        // Nullable options
        this.autoCommit = settings.isAutoCommit();
        this.hasConsumerGroup = settings.getGroupId() != null;

        Map<String, Object> settingsMap = settings.getSettings();
        if (settings.getGroupId() != null) {
            settingsMap.put("group.id", settings.getGroupId() + "-" + GROUP_COUNTER.getAndIncrement());
        }

        this.consumer = new KafkaConsumer<>(settingsMap, new StringDeserializer(), new ByteArrayDeserializer());

        this.scheduler.scheduleAtFixedRate(this::consume, 0, POLL_TIMEOUT_MS, TimeUnit.MILLISECONDS);
        LOGGER.info("Starting Kafka consumer (autoCommit={}, groupId={})", this.autoCommit, settings.getGroupId());
    }

    private void consume() {
        ConsumerRecords<String, byte[]> records;
        synchronized (this.consumedTopics) {
            if (this.consumedTopics.isEmpty()) return;
            records = this.consumer.poll(POLL_TIMEOUT);
            if (this.autoCommit) this.consumer.commitSync();
        }

        for (ConsumerRecord<String, byte[]> record : records) {
            this.processRecord(record);
        }
    }

    private void processRecord(@NotNull ConsumerRecord<String, byte[]> record) {
        String protoType = this.getProtoType(record.headers());
        if (protoType == null) {
            LOGGER.warn("Received message without X-Proto-Type header");
            return;
        }

        AbstractMessage message;
        try {
            message = ProtoParserRegistry.parse(protoType, record.value());
        } catch (InvalidProtocolBufferException exception) {
            LOGGER.warn("Failed to parse message", exception);
            return;
        }

        Set<Consumer<AbstractMessage>> consumers = this.protoListeners.get(message.getClass());
        if (consumers == null) return;

        try {
            for (Consumer<AbstractMessage> consumer : consumers) {
                consumer.accept(message);
            }
        } catch (Exception exception) {
            LOGGER.error("Failed to handle message (topic: {}, type: {}): ", record.topic(), protoType, exception);
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
    public <T extends AbstractMessage> void addListener(@NotNull Class<T> messageType, @NotNull Consumer<T> listener) {
        MessageProtoConfig<T> protoConfig = ProtoParserRegistry.getParser(messageType);
        if (protoConfig == null) {
            throw new IllegalArgumentException("No parser found for " + messageType.getName());
        }

        this.protoListeners.computeIfAbsent(messageType, k -> new HashSet<>()).add((Consumer<AbstractMessage>) listener);

        this.listenToTopic(protoConfig.topic());
    }

    private void listenToTopic(@NotNull String topic) {
        boolean added = this.consumedTopics.add(topic);
        if (!added) return;

        LOGGER.debug("Subscribing to topic {}", topic);
        if (this.hasConsumerGroup) {
            synchronized (this.consumedTopics) {
                this.consumer.subscribe(this.consumedTopics);
            }
        } else {
            TopicPartition partition = new TopicPartition(topic, 0);
            this.partitions.add(partition);
            synchronized (this.consumedTopics) {
                this.consumer.assign(this.partitions);
            }
        }
    }

    public void close() {
        this.scheduler.shutdown();
        synchronized (this.consumedTopics) {
            this.consumer.close();
        }
    }
}
