package dev.emortal.api.utils.callback;

import io.grpc.stub.StreamObserver;
import org.jetbrains.annotations.NotNull;
import org.jetbrains.annotations.Nullable;

import java.util.function.Consumer;

@SuppressWarnings("unused")
public final class FunctionalStreamObserver {

    public static <T> @NotNull StreamObserver<T> create(@NotNull Consumer<T> onNext, @NotNull Consumer<Throwable> onError, @Nullable Runnable onCompleted) {
        return new StreamObserver<>() {
            @Override
            public void onNext(T value) {
                onNext.accept(value);
            }

            @Override
            public void onError(Throwable t) {
                onError.accept(t);
            }

            @Override
            public void onCompleted() {
                if (onCompleted != null) onCompleted.run();
            }
        };
    }

    private FunctionalStreamObserver() {
    }
}
