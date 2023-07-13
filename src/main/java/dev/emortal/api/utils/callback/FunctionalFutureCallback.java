package dev.emortal.api.utils.callback;

import com.google.common.util.concurrent.FutureCallback;
import org.jetbrains.annotations.NotNull;

import java.util.function.Consumer;

@SuppressWarnings("unused")
public final class FunctionalFutureCallback {

    public static <T> @NotNull FutureCallback<T> create(@NotNull Consumer<T> onSuccess, @NotNull Consumer<Throwable> onFailure) {
        return new FutureCallback<>() {
            @Override
            public void onSuccess(T result) {
                onSuccess.accept(result);
            }

            @Override
            public void onFailure(@NotNull Throwable throwable) {
                onFailure.accept(throwable);
            }
        };
    }

    private FunctionalFutureCallback() {
    }
}
