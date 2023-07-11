package dev.emortal.api.service.badges;

import dev.emortal.api.grpc.badge.BadgeManagerGrpc;
import dev.emortal.api.grpc.badge.BadgeManagerProto;
import dev.emortal.api.model.badge.Badge;
import dev.emortal.api.utils.internal.GrpcErrorHelper;
import io.grpc.ManagedChannel;
import io.grpc.StatusRuntimeException;
import org.jetbrains.annotations.ApiStatus;
import org.jetbrains.annotations.NotNull;
import org.jetbrains.annotations.Nullable;

import java.util.List;
import java.util.UUID;

@ApiStatus.Internal
public final class DefaultBadgeService implements BadgeService {

    private final BadgeManagerGrpc.BadgeManagerBlockingStub grpc;

    @ApiStatus.Internal
    public DefaultBadgeService(@NotNull ManagedChannel channel) {
        this.grpc = BadgeManagerGrpc.newBlockingStub(channel);
    }

    @Override
    public @NotNull List<Badge> getAllBadges() {
        var request = BadgeManagerProto.GetBadgesRequest.getDefaultInstance();
        BadgeManagerProto.GetBadgesResponse response = this.grpc.getBadges(request);
        return response.getBadgesList();
    }

    @Override
    public @NotNull PlayerBadges getPlayerBadges(@NotNull UUID playerId) {
        var request = BadgeManagerProto.GetPlayerBadgesRequest.newBuilder().setPlayerId(playerId.toString()).build();
        BadgeManagerProto.GetPlayerBadgesResponse response = this.grpc.getPlayerBadges(request);

        String activeBadgeId = response.getActiveBadgeId();
        return new PlayerBadges(response.getBadgesList(), activeBadgeId.isEmpty() ? null : activeBadgeId);
    }

    @Override
    public @Nullable Badge getActiveBadge(@NotNull UUID playerId) {
        var request = BadgeManagerProto.GetActivePlayerBadgeRequest.newBuilder().setPlayerId(playerId.toString()).build();

        BadgeManagerProto.GetActivePlayerBadgeResponse response;
        try {
            response = this.grpc.getActivePlayerBadge(request);
        } catch (StatusRuntimeException exception) {
            if (exception.getStatus() == io.grpc.Status.NOT_FOUND) return null;
            throw exception;
        }

        return response.getBadge();
    }

    @Override
    public @NotNull SetActiveBadgeResult setActiveBadge(@NotNull UUID playerId, @NotNull String badgeId) {
        var request = BadgeManagerProto.SetActivePlayerBadgeRequest.newBuilder()
                .setPlayerId(playerId.toString())
                .setBadgeId(badgeId)
                .build();

        try {
            this.grpc.setActivePlayerBadge(request);
            return SetActiveBadgeResult.SUCCESS;
        } catch (StatusRuntimeException exception) {
            var error = GrpcErrorHelper.unpackError(exception, BadgeManagerProto.SetActivePlayerBadgeErrorResponse.class);
            if (error.getReason() == BadgeManagerProto.SetActivePlayerBadgeErrorResponse.Reason.PLAYER_DOESNT_HAVE_BADGE) {
                return SetActiveBadgeResult.PLAYER_DOESNT_HAVE_BADGE;
            }
            throw exception;
        }
    }

    @Override
    public @NotNull AddBadgeToPlayerResult addBadgeToPlayer(@NotNull UUID playerId, @NotNull String badgeId) {
        var request = BadgeManagerProto.AddBadgeToPlayerRequest.newBuilder()
                .setPlayerId(playerId.toString())
                .setBadgeId(badgeId)
                .build();

        try {
            this.grpc.addBadgeToPlayer(request);
            return AddBadgeToPlayerResult.SUCCESS;
        } catch (StatusRuntimeException exception) {
            var error = GrpcErrorHelper.unpackError(exception, BadgeManagerProto.AddBadgeToPlayerErrorResponse.class);
            if (error.getReason() == BadgeManagerProto.AddBadgeToPlayerErrorResponse.Reason.PLAYER_ALREADY_HAS_BADGE) {
                return AddBadgeToPlayerResult.PLAYER_ALREADY_HAS_BADGE;
            }
            throw exception;
        }
    }

    @Override
    public @NotNull RemoveBadgeFromPlayerResult removeBadgeFromPlayer(@NotNull UUID playerId, @NotNull String badgeId) {
        var request = BadgeManagerProto.RemoveBadgeFromPlayerRequest.newBuilder()
                .setPlayerId(playerId.toString())
                .setBadgeId(badgeId)
                .build();

        try {
            this.grpc.removeBadgeFromPlayer(request);
            return RemoveBadgeFromPlayerResult.SUCCESS;
        } catch (StatusRuntimeException exception) {
            var error = GrpcErrorHelper.unpackError(exception, BadgeManagerProto.RemoveBadgeFromPlayerErrorResponse.class);
            if (error.getReason() == BadgeManagerProto.RemoveBadgeFromPlayerErrorResponse.Reason.PLAYER_DOESNT_HAVE_BADGE) {
                return RemoveBadgeFromPlayerResult.PLAYER_DOESNT_HAVE_BADGE;
            }
            throw exception;
        }
    }
}
