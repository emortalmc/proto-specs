package dev.emortal.api.service.relationship;

import dev.emortal.api.grpc.relationship.RelationshipGrpc;
import dev.emortal.api.grpc.relationship.RelationshipProto;
import dev.emortal.api.grpc.relationship.RelationshipProto.CreateBlockResponse.CreateBlockResult;
import dev.emortal.api.grpc.relationship.RelationshipProto.DenyFriendRequestResponse.DenyFriendRequestResult;
import dev.emortal.api.grpc.relationship.RelationshipProto.RemoveFriendResponse.RemoveFriendResult;
import dev.emortal.api.model.relationship.FriendRequest;
import dev.emortal.api.model.relationship.PlayerBlock;
import dev.emortal.api.utils.ProtoTimestampConverter;
import io.grpc.Channel;
import io.grpc.Status;
import io.grpc.StatusRuntimeException;
import org.jetbrains.annotations.ApiStatus;
import org.jetbrains.annotations.NotNull;
import org.jetbrains.annotations.Nullable;

import java.time.Instant;
import java.util.ArrayList;
import java.util.List;
import java.util.UUID;

/**
 * The default implementation of {@link RelationshipService} that uses a blocking stub to communicate with the gRPC server.
 */
@ApiStatus.Internal
public final class DefaultRelationshipService implements RelationshipService {

    private final RelationshipGrpc.RelationshipBlockingStub grpc;

    @ApiStatus.Internal
    public DefaultRelationshipService(@NotNull Channel channel) {
        this.grpc = RelationshipGrpc.newBlockingStub(channel);
    }

    @Override
    public @NotNull AddFriendResult addFriend(@NotNull UUID senderId, @NotNull String senderName, @NotNull UUID targetId) {
        var friendRequest = FriendRequest.newBuilder()
                .setSenderId(senderId.toString())
                .setSenderUsername(senderName)
                .setTargetId(targetId.toString())
                .build();
        var request = RelationshipProto.AddFriendRequest.newBuilder().setRequest(friendRequest).build();

        RelationshipProto.AddFriendResponse response = this.grpc.addFriend(request);
        return switch (response.getResult()) {
            case REQUEST_SENT -> new AddFriendResult.RequestSent();
            case FRIEND_ADDED -> new AddFriendResult.FriendAdded(ProtoTimestampConverter.fromProto(response.getFriendsSince()));
            case ALREADY_FRIENDS -> AddFriendResult.Error.ALREADY_FRIENDS;
            case ALREADY_REQUESTED -> AddFriendResult.Error.ALREADY_REQUESTED;
            case YOU_BLOCKED -> AddFriendResult.Error.YOU_BLOCKED;
            case PRIVACY_BLOCKED -> AddFriendResult.Error.PRIVACY_BLOCKED;
            case UNRECOGNIZED -> throw new IllegalStateException("Unrecognized result: " + response.getResult());
        };
    }

    @Override
    public @NotNull RemoveFriendResult removeFriend(@NotNull UUID senderId, @NotNull String senderName, @NotNull UUID targetId) {
        var request = RelationshipProto.RemoveFriendRequest.newBuilder()
                .setSenderId(senderId.toString())
                .setSenderUsername(senderName)
                .setTargetId(targetId.toString())
                .build();

        RelationshipProto.RemoveFriendResponse response = this.grpc.removeFriend(request);
        return response.getResult();
    }

    @Override
    public @NotNull DenyFriendRequestResult denyFriendRequest(@NotNull UUID playerId, @NotNull UUID targetId) {
        var request = RelationshipProto.DenyFriendRequestRequest.newBuilder()
                .setIssuerId(playerId.toString())
                .setTargetId(targetId.toString())
                .build();

        RelationshipProto.DenyFriendRequestResponse response = this.grpc.denyFriendRequest(request);
        return response.getResult();
    }

    @Override
    public int denyAllIncomingFriendRequests(@NotNull UUID playerId) {
        return this.denyAllFriendRequests(playerId, true);
    }

    @Override
    public int denyAllOutgoingFriendRequests(@NotNull UUID playerId) {
        return this.denyAllFriendRequests(playerId, false);
    }

    private int denyAllFriendRequests(@NotNull UUID playerId, boolean incoming) {
        var request = RelationshipProto.MassDenyFriendRequestRequest.newBuilder()
                .setIssuerId(playerId.toString())
                .setIncoming(incoming)
                .build();

        RelationshipProto.MassDenyFriendRequestResponse response = this.grpc.massDenyFriendRequest(request);
        return response.getRequestsDenied();
    }

    @Override
    public @NotNull List<Friend> listFriends(@NotNull UUID playerId) {
        var request = RelationshipProto.GetFriendListRequest.newBuilder().setPlayerId(playerId.toString()).build();
        RelationshipProto.FriendListResponse response = this.grpc.getFriendList(request);

        List<Friend> result = new ArrayList<>();
        for (var entry : response.getFriendsList()) {
            result.add(new Friend(UUID.fromString(entry.getId()), ProtoTimestampConverter.fromProto(entry.getFriendsSince())));
        }

        return result;
    }

    @Override
    public @NotNull List<RequestedFriend> listPendingIncomingFriendRequests(@NotNull UUID playerId) {
        return this.listPendingFriendRequests(playerId, true);
    }

    @Override
    public @NotNull List<RequestedFriend> listPendingOutgoingFriendRequests(@NotNull UUID playerId) {
        return this.listPendingFriendRequests(playerId, false);
    }

    private @NotNull List<RequestedFriend> listPendingFriendRequests(@NotNull UUID playerId, boolean incoming) {
        var request = RelationshipProto.GetPendingFriendRequestListRequest.newBuilder()
                .setIssuerId(playerId.toString())
                .setIncoming(incoming)
                .build();

        RelationshipProto.PendingFriendListResponse response = this.grpc.getPendingFriendRequestList(request);

        List<RequestedFriend> result = new ArrayList<>();
        for (var entry : response.getRequestsList()) {
            UUID requesterId = UUID.fromString(entry.getRequesterId());
            UUID targetId = UUID.fromString(entry.getTargetId());
            Instant requestTime = ProtoTimestampConverter.fromProto(entry.getRequestTime());
            result.add(new RequestedFriend(requesterId, targetId, requestTime));
        }

        return result;
    }

    @Override
    public @NotNull CreateBlockResult block(@NotNull UUID blockerId, @NotNull UUID blockedId) {
        var block = PlayerBlock.newBuilder()
                .setBlockerId(blockerId.toString())
                .setBlockedId(blockedId.toString())
                .build();
        var request = RelationshipProto.CreateBlockRequest.newBuilder().setBlock(block).build();

        RelationshipProto.CreateBlockResponse response = this.grpc.createBlock(request);
        return response.getResult();
    }

    @Override
    public @NotNull DeleteBlockResult unblock(@NotNull UUID blockerId, @NotNull UUID blockedId) {
        var request = RelationshipProto.DeleteBlockRequest.newBuilder()
                .setIssuerId(blockerId.toString())
                .setTargetId(blockedId.toString())
                .build();

        try {
            this.grpc.deleteBlock(request);
            return DeleteBlockResult.SUCCESS;
        } catch (StatusRuntimeException exception) {
            if (exception.getStatus() == Status.NOT_FOUND) return DeleteBlockResult.NOT_BLOCKED;
            throw exception;
        }
    }

    @Override
    public @Nullable PlayerBlock isBlocked(@NotNull UUID blockerId, @NotNull UUID blockedId) {
        var request = RelationshipProto.IsBlockedRequest.newBuilder()
                .setIssuerId(blockerId.toString())
                .setTargetId(blockedId.toString())
                .build();

        RelationshipProto.IsBlockedResponse response = this.grpc.isBlocked(request);
        return response.hasBlock() ? response.getBlock() : null;
    }

    @Override
    public @NotNull List<UUID> getBlockedList(@NotNull UUID blockerId) {
        var request = RelationshipProto.GetBlockedListRequest.newBuilder().setPlayerId(blockerId.toString()).build();
        RelationshipProto.BlockedListResponse response = this.grpc.getBlockedList(request);

        List<UUID> result = new ArrayList<>();
        for (var entry : response.getBlockedPlayerIdsList()) {
            result.add(UUID.fromString(entry));
        }
        return result;
    }
}
