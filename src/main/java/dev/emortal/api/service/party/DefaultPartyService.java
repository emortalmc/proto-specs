package dev.emortal.api.service.party;

import dev.emortal.api.grpc.party.PartyProto;
import dev.emortal.api.grpc.party.PartyServiceGrpc;
import dev.emortal.api.model.common.Pageable;
import dev.emortal.api.model.party.Party;
import dev.emortal.api.model.party.PartyInvite;
import dev.emortal.api.utils.internal.GrpcErrorHelper;
import io.grpc.ManagedChannel;
import io.grpc.Status;
import io.grpc.StatusRuntimeException;
import org.jetbrains.annotations.ApiStatus;
import org.jetbrains.annotations.NotNull;
import org.jetbrains.annotations.Nullable;

import java.util.List;
import java.util.UUID;

@ApiStatus.Internal
public final class DefaultPartyService implements PartyService {

    private final PartyServiceGrpc.PartyServiceBlockingStub grpc;

    @ApiStatus.Internal
    public DefaultPartyService(@NotNull ManagedChannel channel) {
        this.grpc = PartyServiceGrpc.newBlockingStub(channel);
    }

    @Override
    public @Nullable Party getParty(@NotNull String partyId) {
        var request = PartyProto.GetPartyRequest.newBuilder().setPartyId(partyId).build();
        return this.getParty(request);
    }

    @Override
    public @Nullable Party getPartyByPlayer(@NotNull UUID playerId) {
        var request = PartyProto.GetPartyRequest.newBuilder().setPlayerId(playerId.toString()).build();
        return this.getParty(request);
    }

    private @Nullable Party getParty(@NotNull PartyProto.GetPartyRequest request) {
        PartyProto.GetPartyResponse response;
        try {
            response = this.grpc.getParty(request);
        } catch (StatusRuntimeException exception) {
            if (exception.getStatus() == Status.NOT_FOUND) return null;
            throw exception;
        }

        return response.getParty();
    }

    @Override
    public @NotNull ModifyPartyResult emptyParty(@NotNull String partyId) {
        var request = PartyProto.EmptyPartyRequest.newBuilder().setPartyId(partyId).build();
        return this.emptyParty(request);
    }

    @Override
    public @NotNull ModifyPartyResult emptyPartyByPlayer(@NotNull UUID playerId) {
        var request = PartyProto.EmptyPartyRequest.newBuilder().setPlayerId(playerId.toString()).build();
        return this.emptyParty(request);
    }

    private @NotNull ModifyPartyResult emptyParty(@NotNull PartyProto.EmptyPartyRequest request) {
        try {
            this.grpc.emptyParty(request);
            return ModifyPartyResult.SUCCESS;
        } catch (StatusRuntimeException exception) {
            var error = GrpcErrorHelper.unpackError(exception, PartyProto.EmptyPartyErrorResponse.class);
            if (error.getErrorType() == PartyProto.EmptyPartyErrorResponse.ErrorType.NOT_LEADER) return ModifyPartyResult.NOT_LEADER;
            throw exception;
        }
    }

    @Override
    public @NotNull ModifyPartyResult setPartyOpen(@NotNull UUID playerId, boolean open) {
        var request = PartyProto.SetOpenPartyRequest.newBuilder()
                .setPlayerId(playerId.toString())
                .setOpen(open)
                .build();

        try {
            this.grpc.setOpenParty(request);
            return ModifyPartyResult.SUCCESS;
        } catch (StatusRuntimeException exception) {
            var error = GrpcErrorHelper.unpackError(exception, PartyProto.SetOpenPartyErrorResponse.class);
            if (error.getErrorType() == PartyProto.SetOpenPartyErrorResponse.ErrorType.NOT_LEADER) return ModifyPartyResult.NOT_LEADER;
            throw exception;
        }
    }

    @Override
    public @NotNull List<PartyInvite> getInvites(@NotNull String partyId, @NotNull Pageable pageable) {
        var request = PartyProto.GetPartyInvitesRequest.newBuilder()
                .setPartyId(partyId)
                .setPageable(pageable)
                .build();
        return this.getInvites(request);
    }

    @Override
    public @NotNull List<PartyInvite> getInvitesByPlayer(@NotNull UUID playerId, @NotNull Pageable pageable) {
        var request = PartyProto.GetPartyInvitesRequest.newBuilder()
                .setPlayerId(playerId.toString())
                .setPageable(pageable)
                .build();
        return this.getInvites(request);
    }

    private @NotNull List<PartyInvite> getInvites(@NotNull PartyProto.GetPartyInvitesRequest request) {
        PartyProto.GetPartyInvitesResponse response;
        try {
            response = this.grpc.getPartyInvites(request);
        } catch (StatusRuntimeException exception) {
            if (exception.getStatus() == Status.NOT_FOUND) return List.of();
            throw exception;
        }

        return response.getInvitesList();
    }

    @Override
    public @NotNull InvitePlayerToPartyResult invitePlayer(@NotNull UUID inviterId, @NotNull String inviterName,
                                                           @NotNull UUID invitedId, @NotNull String invitedName) {
        var request = PartyProto.InvitePlayerRequest.newBuilder()
                .setIssuerId(inviterId.toString())
                .setIssuerUsername(inviterName)
                .setTargetId(invitedId.toString())
                .setTargetUsername(invitedName)
                .build();

        PartyProto.InvitePlayerResponse response;
        try {
            response = this.grpc.invitePlayer(request);
        } catch (StatusRuntimeException exception) {
            var error = GrpcErrorHelper.unpackError(exception, PartyProto.InvitePlayerErrorResponse.class);

            return switch (error.getErrorType()) {
                case TARGET_ALREADY_IN_SELF_PARTY -> InvitePlayerToPartyResult.Error.TARGET_IN_THIS_PARTY;
                case TARGET_ALREADY_IN_ANOTHER_PARTY -> InvitePlayerToPartyResult.Error.TARGET_IN_OTHER_PARTY;
                case TARGET_ALREADY_INVITED -> InvitePlayerToPartyResult.Error.TARGET_ALREADY_INVITED;
                case NO_PERMISSION -> InvitePlayerToPartyResult.Error.NO_PERMISSION;
                case PARTY_IS_OPEN -> InvitePlayerToPartyResult.Error.PARTY_OPEN;
                case UNRECOGNIZED -> throw exception;
            };
        }

        return new InvitePlayerToPartyResult.Success(response.getInvite());
    }

    @Override
    public @NotNull JoinPartyResult joinParty(@NotNull UUID joinerId, @NotNull String joinerName, @NotNull UUID targetId) {
        var request = PartyProto.JoinPartyRequest.newBuilder()
                .setPlayerId(joinerId.toString())
                .setPlayerUsername(joinerName)
                .setTargetPlayerId(targetId.toString())
                .build();

        PartyProto.JoinPartyResponse response;
        try {
            response = this.grpc.joinParty(request);
        } catch (StatusRuntimeException exception) {
            var error = GrpcErrorHelper.unpackError(exception, PartyProto.JoinPartyErrorResponse.class);

            return switch (error.getErrorType()) {
                case ALREADY_IN_PARTY -> JoinPartyResult.Error.ALREADY_IN_PARTY;
                case NOT_INVITED -> JoinPartyResult.Error.NOT_INVITED;
                case UNRECOGNIZED -> throw exception;
            };
        }

        return new JoinPartyResult.Success(response.getParty());
    }

    @Override
    public @NotNull LeavePartyResult leaveParty(@NotNull UUID leaverId) {
        var request = PartyProto.LeavePartyRequest.newBuilder().setPlayerId(leaverId.toString()).build();

        try {
            this.grpc.leaveParty(request);
            return LeavePartyResult.SUCCESS;
        } catch (StatusRuntimeException exception) {
            var error = GrpcErrorHelper.unpackError(exception, PartyProto.LeavePartyErrorResponse.class);

            return switch (error.getErrorType()) {
                case CANNOT_LEAVE_AS_LEADER -> LeavePartyResult.CANNOT_LEAVE_AS_LEADER;
                case UNRECOGNIZED -> throw exception;
            };
        }
    }

    @Override
    public @NotNull KickPlayerFromPartyResult kickPlayer(@NotNull UUID kickerId, @NotNull String kickerName, @NotNull UUID targetId) {
        var request = PartyProto.KickPlayerRequest.newBuilder()
                .setIssuerId(kickerId.toString())
                .setIssuerUsername(kickerName)
                .setTargetId(targetId.toString())
                .build();

        try {
            this.grpc.kickPlayer(request);
            return KickPlayerFromPartyResult.SUCCESS;
        } catch (StatusRuntimeException exception) {
            var error = GrpcErrorHelper.unpackError(exception, PartyProto.KickPlayerErrorResponse.class);

            return switch (error.getErrorType()) {
                case TARGET_NOT_IN_PARTY -> KickPlayerFromPartyResult.TARGET_NOT_IN_PARTY;
                case SELF_NOT_LEADER -> KickPlayerFromPartyResult.SELF_NOT_LEADER;
                case TARGET_IS_LEADER -> KickPlayerFromPartyResult.TARGET_IS_LEADER;
                case UNRECOGNIZED -> throw exception;
            };
        }
    }

    @Override
    public @NotNull SetPartyLeaderResult setPartyLeader(@NotNull UUID requesterId, @NotNull String requesterName, @NotNull UUID newLeaderId) {
        var request = PartyProto.SetPartyLeaderRequest.newBuilder()
                .setIssuerId(requesterId.toString())
                .setIssuerUsername(requesterName)
                .setTargetId(newLeaderId.toString())
                .build();

        try {
            this.grpc.setPartyLeader(request);
            return SetPartyLeaderResult.SUCCESS;
        } catch (StatusRuntimeException exception) {
            var error = GrpcErrorHelper.unpackError(exception, PartyProto.SetPartyLeaderErrorResponse.class);

            return switch (error.getErrorType()) {
                case TARGET_NOT_IN_PARTY -> SetPartyLeaderResult.TARGET_NOT_IN_PARTY;
                case SELF_NOT_LEADER -> SetPartyLeaderResult.SELF_NOT_LEADER;
                case UNRECOGNIZED -> throw exception;
            };
        }
    }
}
