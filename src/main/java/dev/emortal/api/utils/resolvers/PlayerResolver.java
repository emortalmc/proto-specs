package dev.emortal.api.utils.resolvers;

import dev.emortal.api.model.mcplayer.McPlayer;
import dev.emortal.api.service.mcplayer.McPlayerService;
import dev.emortal.api.utils.GrpcStubCollection;
import io.grpc.Status;
import io.grpc.StatusRuntimeException;

import java.util.UUID;
import java.util.function.Consumer;
import java.util.function.Function;

@SuppressWarnings("unused")
public class PlayerResolver {
    private static final McPlayerService playerService = GrpcStubCollection.getPlayerService().orElse(null);
    // The alternative should be used as a server software specific option to avoid calling the mc-player service if they're on the same server.
    private static Function<String, CachedMcPlayer> platformUsernameResolver;

    public static void setPlatformUsernameResolver(Function<String, CachedMcPlayer> platformUsernameResolver) {
        if (PlayerResolver.platformUsernameResolver != null) throw new IllegalStateException("Player service already set");
        PlayerResolver.platformUsernameResolver = platformUsernameResolver;
    }

    public static void retrievePlayerData(String username, Consumer<CachedMcPlayer> callback, Consumer<Status> errorCallback) {
        String usernameLowercase = username.toLowerCase();

        CachedMcPlayer alternativeOption = platformUsernameResolver.apply(usernameLowercase);
        if (alternativeOption != null) {
            callback.accept(alternativeOption);
            return;
        }

        Thread.startVirtualThread(() -> requestMcPlayer(usernameLowercase, callback, errorCallback));
    }

    private static void requestMcPlayer(String username, Consumer<CachedMcPlayer> callback, Consumer<Status> errorCallback) {
        McPlayer player;
        try {
            player = playerService.getPlayerByUsername(username);
        } catch (StatusRuntimeException exception) {
            errorCallback.accept(exception.getStatus());
            return;
        }

        if (player == null) return;
        callback.accept(new CachedMcPlayer(UUID.fromString(player.getId()), player.getCurrentUsername(), player.hasCurrentServer()));
    }

    public record CachedMcPlayer(UUID uuid, String username, boolean online) {}
}
