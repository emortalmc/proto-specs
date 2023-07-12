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
import org.apache.kafka.clients.producer.RecordMetadata;
import org.apache.kafka.common.header.Header;
import org.apache.kafka.common.header.internals.RecordHeader;
import org.apache.kafka.common.serialization.ByteArraySerializer;
import org.apache.kafka.common.serialization.StringSerializer;
import org.jetbrains.annotations.CheckReturnValue;
import org.jetbrains.annotations.NotNull;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.net.InetAddress;
import java.net.UnknownHostException;
import java.util.List;
import java.util.concurrent.ForkJoinPool;
import java.util.concurrent.Future;

public final class FriendlyKafkaProducer {
    private static final Logger LOGGER = LoggerFactory.getLogger(FriendlyKafkaProducer.class);

    private static final String HOSTNAME = getHostname();

    private final @NotNull KafkaProducer<String, byte[]> producer;

    public FriendlyKafkaProducer(@NotNull KafkaSettings settings) {
        this.producer = new KafkaProducer<>(settings.getSettings(), new StringSerializer(), new ByteArraySerializer());
    }

    @CheckReturnValue
    public @NotNull ListenableFuture<Void> produce(@NotNull String topic, @NotNull AbstractMessage value) {
        List<Header> headers = List.of(
                new RecordHeader("X-Source-Id", HOSTNAME.getBytes()),
                new RecordHeader("X-Proto-Type", value.getDescriptorForType().getFullName().getBytes())
        );
        ProducerRecord<String, byte[]> record = new ProducerRecord<>(topic, null, System.currentTimeMillis(), null, value.toByteArray(), headers);

        Future<RecordMetadata> future = this.producer.send(record);

        return Futures.transform(
                JdkFutureAdapters.listenInPoolThread(future, ForkJoinPool.commonPool()),
                unused -> null, ForkJoinPool.commonPool()
        );
    }

    @CheckReturnValue
    public @NotNull ListenableFuture<Void> produce(@NotNull AbstractMessage value) {
        MessageProtoConfig<?> config = ProtoParserRegistry.getParser(value.getClass());
        if (config == null) throw new IllegalArgumentException("No parser found for message type " + value.getClass().getName());

        return this.produce(config.topic(), value);
    }

    public void produceAndForget(@NotNull String topic, @NotNull AbstractMessage value) {
        ListenableFuture<Void> future = this.produce(topic, value);

        Futures.addCallback(future, FunctionalFutureCallback.create(
                unused -> {},
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

    private static @NotNull String getHostname() {
        try {
            return InetAddress.getLocalHost().getHostName();
        } catch (UnknownHostException exception) {
            LOGGER.error("Failed to get hostname", exception);
            return "unknown";
        }
    }
}