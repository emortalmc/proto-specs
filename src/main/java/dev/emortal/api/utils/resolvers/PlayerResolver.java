package dev.emortal.api.utils.resolvers;

import dev.emortal.api.model.mcplayer.McPlayer;
import dev.emortal.api.service.mcplayer.McPlayerService;
import dev.emortal.api.utils.GrpcStubCollection;
import io.grpc.Status;
import io.grpc.StatusException;
import io.grpc.StatusRuntimeException;
import org.jetbrains.annotations.Blocking;
import org.jetbrains.annotations.NonBlocking;
import org.jetbrains.annotations.NotNull;
import org.jetbrains.annotations.Nullable;

import java.util.Locale;
import java.util.UUID;
import java.util.function.Consumer;
import java.util.function.Function;

/**
 * A utility for resolving player data from a local resolver, or the mc player service, if the player is not available locally.
 */
@SuppressWarnings("unused")
public final class PlayerResolver {
    private static final McPlayerService playerService = GrpcStubCollection.getPlayerService().orElse(null);
    // The alternative should be used as a server software specific option to avoid calling the mc-player service if they're on the same server.
    private static @Nullable Function<String, CachedMcPlayer> platformUsernameResolver;

    public static void setPlatformUsernameResolver(@NotNull Function<String, CachedMcPlayer> platformUsernameResolver) {
        if (PlayerResolver.platformUsernameResolver != null) throw new IllegalStateException("Platform resolver already set!");
        PlayerResolver.platformUsernameResolver = platformUsernameResolver;
    }

    /**
     * This will first check the platform resolver, and if it returns null, it will check the mc player service.
     */
    @Blocking
    public static @Nullable CachedMcPlayer getPlayerData(@NotNull String username) throws StatusException {
        String usernameLowercase = username.toLowerCase(Locale.ROOT);

        CachedMcPlayer alternativeOption = platformUsernameResolver.apply(usernameLowercase);
        if (alternativeOption != null) return alternativeOption;

        try {
            return requestMcPlayer(usernameLowercase);
        } catch (StatusRuntimeException exception) {
            throw new StatusException(exception.getStatus(), exception.getTrailers());
        }
    }

    /**
     * This will first check the platform resolver, and if it returns null, it will check the mc player service.
     *
     * <p>
     * This method is asynchronous, and will offload the request to the mc player service to a virtual thread.
     */
    @NonBlocking
    public static void retrievePlayerData(@NotNull String username, @NotNull Consumer<CachedMcPlayer> callback, @NotNull Consumer<Status> errorCallback) {
        String usernameLowercase = username.toLowerCase();

        CachedMcPlayer alternativeOption = platformUsernameResolver.apply(usernameLowercase);
        if (alternativeOption != null) {
            callback.accept(alternativeOption);
            return;
        }

        Thread.startVirtualThread(() -> {
            try {
                CachedMcPlayer player = requestMcPlayer(usernameLowercase);
                callback.accept(player);
            } catch (StatusRuntimeException exception) {
                errorCallback.accept(exception.getStatus());
            }
        });
    }

    @Blocking
    private static @Nullable CachedMcPlayer requestMcPlayer(@NotNull String username) throws StatusRuntimeException {
        McPlayer player = playerService.getPlayerByUsername(username);
        if (player == null) return null;

        return new CachedMcPlayer(UUID.fromString(player.getId()), player.getCurrentUsername(), player.hasCurrentServer());
    }

    private PlayerResolver() {
    }

    public record CachedMcPlayer(UUID uuid, String username, boolean online) {
    }
}
