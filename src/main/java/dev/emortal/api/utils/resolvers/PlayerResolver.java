package dev.emortal.api.utils.resolvers;

import dev.emortal.api.model.mcplayer.McPlayer;
import dev.emortal.api.service.mcplayer.McPlayerService;
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

    private final McPlayerService playerService;
    // The platform resolver is used to resolve a player's data from a local source to avoid calling the service if they're on the same server.
    private final Function<String, CachedMcPlayer> platformResolver;

    public PlayerResolver(@NotNull McPlayerService playerService, @NotNull Function<String, CachedMcPlayer> platformResolver) {
        this.playerService = playerService;
        this.platformResolver = platformResolver;
    }

    /**
     * This will first check the platform resolver, and if it returns null, it will check the mc player service.
     */
    @Blocking
    public @Nullable CachedMcPlayer getPlayerData(@NotNull String username) throws StatusException {
        String usernameLowercase = username.toLowerCase(Locale.ROOT);

        CachedMcPlayer alternativeOption = this.platformResolver.apply(usernameLowercase);
        if (alternativeOption != null) return alternativeOption;

        try {
            return this.requestMcPlayer(usernameLowercase);
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
    public void retrievePlayerData(@NotNull String username, @NotNull Consumer<CachedMcPlayer> callback, @NotNull Consumer<Status> errorCallback) {
        String usernameLowercase = username.toLowerCase();

        CachedMcPlayer alternativeOption = this.platformResolver.apply(usernameLowercase);
        if (alternativeOption != null) {
            callback.accept(alternativeOption);
            return;
        }

        Thread.startVirtualThread(() -> {
            try {
                CachedMcPlayer player = this.requestMcPlayer(usernameLowercase);
                callback.accept(player);
            } catch (StatusRuntimeException exception) {
                errorCallback.accept(exception.getStatus());
            }
        });
    }

    @Blocking
    private @Nullable CachedMcPlayer requestMcPlayer(@NotNull String username) throws StatusRuntimeException {
        McPlayer player = this.playerService.getPlayerByUsername(username);
        if (player == null) return null;

        return new CachedMcPlayer(UUID.fromString(player.getId()), player.getCurrentUsername(), player.hasCurrentServer());
    }

    public record CachedMcPlayer(@NotNull UUID uuid, @NotNull String username, boolean online) {
    }
}
