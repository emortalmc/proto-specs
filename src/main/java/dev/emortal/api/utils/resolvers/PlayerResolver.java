package dev.emortal.api.utils.resolvers;

import com.google.common.util.concurrent.Futures;
import dev.emortal.api.grpc.mcplayer.McPlayerGrpc;
import dev.emortal.api.grpc.mcplayer.McPlayerProto;
import dev.emortal.api.model.mcplayer.McPlayer;
import dev.emortal.api.utils.GrpcStubCollection;
import dev.emortal.api.utils.callback.FunctionalFutureCallback;
import io.grpc.Status;

import java.util.UUID;
import java.util.concurrent.ForkJoinPool;
import java.util.function.Consumer;
import java.util.function.Function;

@SuppressWarnings("unused")
public class PlayerResolver {
    private final static McPlayerGrpc.McPlayerFutureStub PLAYER_SERVICE = GrpcStubCollection.getPlayerService().orElse(null);
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

        requestMcPlayer(usernameLowercase, callback, errorCallback);
    }

    private static void requestMcPlayer(String username, Consumer<CachedMcPlayer> callback, Consumer<Status> errorCallback) {
        var playerResponseFuture = PLAYER_SERVICE
                .getPlayerByUsername(McPlayerProto.PlayerUsernameRequest.newBuilder().setUsername(username).build());

        Futures.addCallback(playerResponseFuture, FunctionalFutureCallback.create(
                response -> {
                    McPlayer player = response.getPlayer();
                    callback.accept(new CachedMcPlayer(UUID.fromString(player.getId()), player.getCurrentUsername(), player.getCurrentlyOnline()));
                },
                throwable -> errorCallback.accept(Status.fromThrowable(throwable))
        ), ForkJoinPool.commonPool());
    }

    public record CachedMcPlayer(UUID uuid, String username, boolean online) {}
}
