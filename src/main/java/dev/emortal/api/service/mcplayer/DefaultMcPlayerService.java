package dev.emortal.api.service.mcplayer;

import dev.emortal.api.grpc.mcplayer.McPlayerGrpc;
import dev.emortal.api.grpc.mcplayer.McPlayerProto;
import dev.emortal.api.grpc.mcplayer.McPlayerProto.SearchPlayersByUsernameRequest.FilterMethod;
import dev.emortal.api.model.common.Pageable;
import dev.emortal.api.model.mcplayer.McPlayer;
import io.grpc.Channel;
import io.grpc.Status;
import io.grpc.StatusRuntimeException;
import org.jetbrains.annotations.ApiStatus;
import org.jetbrains.annotations.NotNull;
import org.jetbrains.annotations.Nullable;

import java.util.Collection;
import java.util.List;
import java.util.UUID;

/**
 * The default implementation of {@link McPlayerService} that uses a blocking stub to communicate with the gRPC server.
 */
@ApiStatus.Internal
public final class DefaultMcPlayerService implements McPlayerService {

    private final McPlayerGrpc.McPlayerBlockingStub grpc;

    @ApiStatus.Internal
    public DefaultMcPlayerService(@NotNull Channel channel) {
        this.grpc = McPlayerGrpc.newBlockingStub(channel);
    }

    @Override
    public @Nullable McPlayer getPlayerById(@NotNull UUID id) {
        var request = McPlayerProto.GetPlayerRequest.newBuilder().setPlayerId(id.toString()).build();

        McPlayerProto.GetPlayerResponse response;
        try {
            response = this.grpc.getPlayer(request);
        } catch (StatusRuntimeException exception) {
            if (exception.getStatus() == Status.NOT_FOUND) return null;
            throw exception;
        }

        return response.getPlayer();
    }

    @Override
    public @NotNull List<McPlayer> getPlayersById(@NotNull Collection<UUID> ids) {
        var requestBuilder = McPlayerProto.GetPlayersRequest.newBuilder();
        for (UUID id : ids) {
            requestBuilder.addPlayerIds(id.toString());
        }
        var request = requestBuilder.build();

        McPlayerProto.GetPlayersResponse response = this.grpc.getPlayers(request);
        return response.getPlayersList();
    }

    @Override
    public @Nullable McPlayer getPlayerByUsername(@NotNull String username) {
        var request = McPlayerProto.PlayerUsernameRequest.newBuilder().setUsername(username).build();

        McPlayerProto.GetPlayerByUsernameResponse response;
        try {
            response = this.grpc.getPlayerByUsername(request);
        } catch (StatusRuntimeException exception) {
            if (exception.getStatus() == Status.NOT_FOUND) return null;
            throw exception;
        }

        return response.getPlayer();
    }

    @Override
    public @NotNull List<McPlayer> searchPlayersByUsername(@NotNull UUID requesterId, @NotNull String searchUsername, @NotNull Pageable pageable,
                                                           @Nullable FilterMethod filterMethod) {
        var requestBuilder = McPlayerProto.SearchPlayersByUsernameRequest.newBuilder()
                .setIssuerId(requesterId.toString())
                .setSearchUsername(searchUsername)
                .setPageable(pageable);
        if (filterMethod != null) requestBuilder.setFilterMethod(filterMethod);
        var request = requestBuilder.build();

        McPlayerProto.SearchPlayersByUsernameResponse response = this.grpc.searchPlayersByUsername(request);
        return response.getPlayersList();
    }

    @Override
    public @NotNull LoginSessions getLoginSessions(@NotNull UUID playerId, @NotNull Pageable pageable) {
        var request = McPlayerProto.GetLoginSessionsRequest.newBuilder()
                .setPlayerId(playerId.toString())
                .setPageable(pageable)
                .build();

        McPlayerProto.LoginSessionsResponse response = this.grpc.getLoginSessions(request);
        return new LoginSessions(response.getSessionsList(), response.getPageData());
    }
}
