package dev.emortal.api.service.matchmaker;

import dev.emortal.api.grpc.matchmaker.MatchmakerGrpc;
import dev.emortal.api.grpc.matchmaker.MatchmakerProto;
import dev.emortal.api.utils.internal.GrpcErrorHelper;
import io.grpc.Channel;
import io.grpc.StatusRuntimeException;
import org.jetbrains.annotations.ApiStatus;
import org.jetbrains.annotations.NotNull;
import org.jetbrains.annotations.Nullable;

import java.util.Collection;
import java.util.UUID;

@ApiStatus.Internal
public final class DefaultMatchmakerService implements MatchmakerService {

    private final MatchmakerGrpc.MatchmakerBlockingStub grpc;

    @ApiStatus.Internal
    public DefaultMatchmakerService(@NotNull Channel channel) {
        this.grpc = MatchmakerGrpc.newBlockingStub(channel);
    }

    @Override
    public @Nullable MatchmakerProto.GetPlayerQueueInfoResponse getPlayerQueueInfo(@NotNull UUID playerId) {
        var request = MatchmakerProto.GetPlayerQueueInfoRequest.newBuilder().setPlayerId(playerId.toString()).build();

        try {
            return this.grpc.getPlayerQueueInfo(request);
        } catch (StatusRuntimeException exception) {
            if (exception.getStatus().getCode() == io.grpc.Status.Code.NOT_FOUND) return null;
            throw exception;
        }
    }

    @Override
    public @NotNull QueuePlayerResult queuePlayer(@NotNull String gameModeId, @NotNull UUID playerId, @NotNull QueueOptions options) {
        var requestBuilder = MatchmakerProto.QueueByPlayerRequest.newBuilder().setGameModeId(gameModeId).setPlayerId(playerId.toString());
        if (options.privateGame()) requestBuilder.setPrivateGame(true);
        if (options.mapId() != null) requestBuilder.setMapId(options.mapId());
        if (options.autoTeleport()) requestBuilder.setAutoTeleport(true);
        var request = requestBuilder.build();

        try {
            this.grpc.queueByPlayer(request);
            return QueuePlayerResult.SUCCESS;
        } catch (StatusRuntimeException exception) {
            var error = GrpcErrorHelper.unpackError(exception, MatchmakerProto.QueueByPlayerErrorResponse.class);

            return switch (error.getReason()) {
                case ALREADY_IN_QUEUE -> QueuePlayerResult.ALREADY_IN_QUEUE;
                case INVALID_GAME_MODE -> QueuePlayerResult.INVALID_GAME_MODE;
                case GAME_MODE_DISABLED -> QueuePlayerResult.GAME_MODE_DISABLED;
                case INVALID_MAP -> QueuePlayerResult.INVALID_MAP;
                case PARTY_TOO_LARGE -> QueuePlayerResult.PARTY_TOO_LARGE;
                case NO_PERMISSION -> QueuePlayerResult.NO_PERMISSION;
                case UNRECOGNIZED -> throw exception;
            };
        }
    }

    @Override
    public @NotNull DequeuePlayerResult dequeuePlayer(@NotNull UUID playerId) {
        var request = MatchmakerProto.DequeueByPlayerRequest.newBuilder().setPlayerId(playerId.toString()).build();

        try {
            this.grpc.dequeueByPlayer(request);
            return DequeuePlayerResult.SUCCESS;
        } catch (StatusRuntimeException exception) {
            var error = GrpcErrorHelper.unpackError(exception, MatchmakerProto.DequeueByPlayerErrorResponse.class);

            return switch (error.getReason()) {
                case NOT_IN_QUEUE -> DequeuePlayerResult.NOT_IN_QUEUE;
                case NO_PERMISSION -> DequeuePlayerResult.NO_PERMISSION;
                case ALREADY_MARKED_FOR_DEQUEUE -> DequeuePlayerResult.ALREADY_MARKED_FOR_DEQUEUE;
                case UNRECOGNIZED -> throw exception;
            };
        }
    }

    @Override
    public void sendPlayerToLobby(@NotNull UUID playerId, boolean sendParty) {
        var request = MatchmakerProto.SendPlayerToLobbyRequest.newBuilder()
                .addPlayerIds(playerId.toString())
                .setSendParties(sendParty)
                .build();
        this.grpc.sendPlayersToLobby(request);
    }

    @Override
    public void sendPlayersToLobby(@NotNull Collection<UUID> playerIds, boolean sendParties) {
        var requestBuilder = MatchmakerProto.SendPlayerToLobbyRequest.newBuilder();
        for (UUID id : playerIds) {
            requestBuilder.addPlayerIds(id.toString());
        }
        var request = requestBuilder.setSendParties(sendParties).build();

        this.grpc.sendPlayersToLobby(request);
    }

    @Override
    public void loginQueue(@NotNull UUID playerId, boolean proxy) {
        var request = MatchmakerProto.LoginQueueByPlayerRequest.newBuilder().setPlayerId(playerId.toString())
                .setIsProxy(proxy).build();
        this.grpc.loginQueueByPlayer(request);
    }

    @Override
    public @NotNull ChangeMapVoteResult changeMapVote(@NotNull UUID playerId, @NotNull String newMapId) {
        var request = MatchmakerProto.ChangePlayerMapVoteRequest.newBuilder()
                .setPlayerId(playerId.toString())
                .setMapId(newMapId)
                .build();

        try {
            this.grpc.changePlayerMapVote(request);
            return ChangeMapVoteResult.SUCCESS;
        } catch (StatusRuntimeException exception) {
            var error = GrpcErrorHelper.unpackError(exception, MatchmakerProto.ChangePlayerMapVoteErrorResponse.class);

            return switch (error.getReason()) {
                case NOT_IN_QUEUE -> ChangeMapVoteResult.NOT_IN_QUEUE;
                case INVALID_MAP -> ChangeMapVoteResult.INVALID_MAP;
                case UNRECOGNIZED -> throw exception;
            };
        }
    }
}
