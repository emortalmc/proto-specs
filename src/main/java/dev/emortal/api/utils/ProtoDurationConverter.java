package dev.emortal.api.utils;

import com.google.protobuf.Duration;
import com.google.protobuf.Timestamp;
import org.jetbrains.annotations.NotNull;

import java.time.Instant;

public final class ProtoDurationConverter {
    private static final long NANOS_PER_SECOND = 1_000_000_000;

    public static @NotNull Duration toProto(@NotNull java.time.Duration duration) {
        return Duration.newBuilder()
                .setSeconds(duration.getSeconds())
                .setNanos(duration.getNano())
                .build();
    }

    public static @NotNull java.time.Duration fromProto(@NotNull Duration duration) {
        return java.time.Duration.ofSeconds(duration.getSeconds(), duration.getNanos());
    }

    public static @NotNull Duration toProtoFromMillis(long millis) {
        return Duration.newBuilder()
                .setSeconds(millis / 1000)
                .setNanos((int) ((millis % 1000) * 1_000_000))
                .build();
    }

    public static @NotNull Duration toProtoFromNanos(long nanos) {
        return Duration.newBuilder()
                .setSeconds(nanos / NANOS_PER_SECOND)
                .setNanos((int) (nanos % NANOS_PER_SECOND))
                .build();
    }

    public static @NotNull java.time.Duration betweenFromProto(@NotNull Timestamp start, @NotNull Timestamp end) {
        return java.time.Duration.between(ProtoTimestampConverter.fromProto(start), ProtoTimestampConverter.fromProto(end));
    }

    public static @NotNull Duration betweenToProto(@NotNull Instant start, @NotNull Instant end) {
        return toProto(java.time.Duration.between(start, end));
    }

    private ProtoDurationConverter() {
    }
}
