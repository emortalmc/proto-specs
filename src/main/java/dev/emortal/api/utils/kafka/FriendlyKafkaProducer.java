package dev.emortal.api.utils.kafka;

import com.google.common.util.concurrent.Futures;
import com.google.common.util.concurrent.JdkFutureAdapters;
import com.google.common.util.concurrent.ListenableFuture;
import com.google.protobuf.AbstractMessage;
import dev.emortal.api.utils.callback.FunctionalFutureCallback;
import dev.emortal.api.utils.parser.MessageProtoConfig;
import dev.emortal.api.utils.parser.ProtoParserRegistry;
import org.apache.kafka.clients.producer.KafkaProducer;
import org.apache.kafka.clients.producer.ProducerRecord;
import org.apache.kafka.common.header.internals.RecordHeader;
import org.apache.kafka.common.serialization.ByteArraySerializer;
import org.apache.kafka.common.serialization.StringSerializer;
import org.jetbrains.annotations.NotNull;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.util.List;
import java.util.concurrent.ForkJoinPool;

public class FriendlyKafkaProducer {
    private static final Logger LOGGER = LoggerFactory.getLogger(FriendlyKafkaProducer.class);

    private final KafkaProducer<String, byte[]> producer;

    public FriendlyKafkaProducer(@NotNull KafkaSettings settings) {
        this.producer = new KafkaProducer<>(settings.getSettings(), new StringSerializer(), new ByteArraySerializer());
    }

    public @NotNull ListenableFuture<Void> produce(@NotNull String topic, @NotNull AbstractMessage value) {
        ProducerRecord<String, byte[]> record = new ProducerRecord<>(topic, null, System.currentTimeMillis(),
                null, value.toByteArray(), List.of(
                new RecordHeader("X-Proto-Type", value.getDescriptorForType().getFullName().getBytes())
        ));

        var future = this.producer.send(record);

        return Futures.transform(
                JdkFutureAdapters.listenInPoolThread(future, ForkJoinPool.commonPool()),
                unused -> null, ForkJoinPool.commonPool()
        );
    }

    public @NotNull ListenableFuture<Void> produce(@NotNull AbstractMessage value) {
        MessageProtoConfig<?> config = ProtoParserRegistry.getParser(value.getClass());
        if (config == null) throw new IllegalArgumentException("No parser found for message type " + value.getClass().getName());

        return this.produce(config.topic(), value);
    }

    public void produceAndForget(@NotNull String topic, @NotNull AbstractMessage value) {
        ProducerRecord<String, byte[]> record = new ProducerRecord<>(topic, null, System.currentTimeMillis(),
                null, value.toByteArray(), List.of(
                new RecordHeader("X-Proto-Type", value.getDescriptorForType().getFullName().getBytes())
        ));

        var future = JdkFutureAdapters.listenInPoolThread(this.producer.send(record), ForkJoinPool.commonPool());

        Futures.addCallback(future, FunctionalFutureCallback.create(
                unused -> {
                },
                throwable -> LOGGER.error("Failed to produce message to topic " + topic, throwable)
        ), Runnable::run);
    }

    public void produceAndForget(@NotNull AbstractMessage value) {
        MessageProtoConfig<?> config = ProtoParserRegistry.getParser(value.getClass());
        if (config == null) throw new IllegalArgumentException("No parser found for message type " + value.getClass().getName());

        this.produceAndForget(config.topic(), value);
    }

    public void shutdown() {
        this.producer.close();
    }
}