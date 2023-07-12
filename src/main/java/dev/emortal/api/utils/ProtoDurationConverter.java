package dev.emortal.api.utils;

import com.google.protobuf.Duration;
import org.jetbrains.annotations.NotNull;

public final class ProtoDurationConverter {

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
                .setNanos((int) ((millis % 1000) * 1000000))
                .build();
    }

    public static @NotNull Duration toProtoFromNanos(long nanos) {
        return Duration.newBuilder()
                .setSeconds(nanos / 1000000000)
                .setNanos((int) (nanos % 1000000000))
                .build();
    }

    private ProtoDurationConverter() {
    }
}
