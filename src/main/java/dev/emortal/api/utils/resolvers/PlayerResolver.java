package dev.emortal.api.utils.resolvers;

import dev.emortal.api.model.mcplayer.McPlayer;
import dev.emortal.api.service.mcplayer.McPlayerService;
import io.grpc.StatusException;
import io.grpc.StatusRuntimeException;
import org.jetbrains.annotations.Blocking;
import org.jetbrains.annotations.NotNull;
import org.jetbrains.annotations.Nullable;

import java.util.Locale;
import java.util.UUID;
import java.util.function.Function;

/**
 * A utility for resolving player data from a local resolver, or the mc player service, if the player is not available locally.
 */
@SuppressWarnings("unused")
public final class PlayerResolver {

    private final McPlayerService playerService;
    private final Function<UUID, CachedMcPlayer> platformResolverById;
    // The platform resolver is used to resolve a player's data from a local source to avoid calling the service if they're on the same server.
    private final Function<String, CachedMcPlayer> platformResolverByName;

    public PlayerResolver(@NotNull McPlayerService playerService, @NotNull Function<UUID, CachedMcPlayer> platformResolverById,
                          @NotNull Function<String, CachedMcPlayer> platformResolverByName) {
        this.playerService = playerService;
        this.platformResolverById = platformResolverById;
        this.platformResolverByName = platformResolverByName;
    }

    @Blocking
    public @Nullable CachedMcPlayer getPlayerData(@NotNull UUID uuid) throws StatusException {
        CachedMcPlayer platformPlayer = this.platformResolverById.apply(uuid);
        if (platformPlayer != null) return platformPlayer;

        return this.requestMcPlayer(uuid);
    }

    @Blocking
    public @Nullable CachedMcPlayer getPlayerData(@NotNull String username) throws StatusException {
        String usernameLowercase = username.toLowerCase(Locale.ROOT);

        CachedMcPlayer platformPlayer = this.platformResolverByName.apply(usernameLowercase);
        if (platformPlayer != null) return platformPlayer;

        return this.requestMcPlayer(usernameLowercase);
    }

    @Blocking
    private @Nullable CachedMcPlayer requestMcPlayer(@NotNull UUID uuid) throws StatusException {
        McPlayer player;
        try {
            player = this.playerService.getPlayerById(uuid);
        } catch (StatusRuntimeException exception) {
            throw new StatusException(exception.getStatus(), exception.getTrailers());
        }

        if (player == null) return null;
        return new CachedMcPlayer(UUID.fromString(player.getId()), player.getCurrentUsername(), player.hasCurrentServer());
    }

    @Blocking
    private @Nullable CachedMcPlayer requestMcPlayer(@NotNull String username) throws StatusException {
        McPlayer player;
        try {
            player = this.playerService.getPlayerByUsername(username);
        } catch (StatusRuntimeException exception) {
            throw new StatusException(exception.getStatus(), exception.getTrailers());
        }

        if (player == null) return null;
        return new CachedMcPlayer(UUID.fromString(player.getId()), player.getCurrentUsername(), player.hasCurrentServer());
    }

    public record CachedMcPlayer(@NotNull UUID uuid, @NotNull String username, boolean online) {
    }
}
