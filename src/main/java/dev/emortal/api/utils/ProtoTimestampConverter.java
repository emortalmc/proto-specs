package dev.emortal.api.utils;

import com.google.protobuf.Timestamp;
import org.jetbrains.annotations.NotNull;

import java.time.Instant;

public final class ProtoTimestampConverter {
    private static final long NANOS_PER_SECOND = 1_000_000_000;

    public static @NotNull Timestamp toProto(@NotNull Instant instant) {
        return Timestamp.newBuilder()
                .setSeconds(instant.getEpochSecond())
                .setNanos(instant.getNano())
                .build();
    }

    public static @NotNull Instant fromProto(@NotNull Timestamp timestamp) {
        return Instant.ofEpochSecond(timestamp.getSeconds(), timestamp.getNanos());
    }

    public static @NotNull Timestamp now() {
        return toProto(Instant.now());
    }

    public static @NotNull Timestamp toProtoFromMillis(long millis) {
        if (millis < 0) throw new IllegalArgumentException("Millis cannot be negative");

        return Timestamp.newBuilder()
                .setSeconds(millis / 1000)
                .setNanos((int) ((millis % 1000) * 1_000_000))
                .build();
    }

    public static @NotNull Timestamp toProtoFromNanos(long nanos) {
        if (nanos < 0) throw new IllegalArgumentException("Nanos cannot be negative");

        return Timestamp.newBuilder()
                .setSeconds(nanos / NANOS_PER_SECOND)
                .setNanos((int) (nanos % NANOS_PER_SECOND))
                .build();
    }

    private ProtoTimestampConverter() {
    }
}
