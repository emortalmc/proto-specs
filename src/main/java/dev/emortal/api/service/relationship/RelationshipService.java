package dev.emortal.api.service.relationship;

import dev.emortal.api.grpc.relationship.RelationshipProto.CreateBlockResponse.CreateBlockResult;
import dev.emortal.api.grpc.relationship.RelationshipProto.DenyFriendRequestResponse.DenyFriendRequestResult;
import dev.emortal.api.grpc.relationship.RelationshipProto.RemoveFriendResponse.RemoveFriendResult;
import dev.emortal.api.model.relationship.PlayerBlock;
import org.jetbrains.annotations.NotNull;
import org.jetbrains.annotations.Nullable;

import java.util.List;
import java.util.UUID;

public interface RelationshipService {

    @NotNull AddFriendResult addFriend(@NotNull UUID senderId, @NotNull String senderName, @NotNull UUID targetId);

    @NotNull RemoveFriendResult removeFriend(@NotNull UUID senderId, @NotNull String senderName, @NotNull UUID targetId);

    @NotNull DenyFriendRequestResult denyFriendRequest(@NotNull UUID playerId, @NotNull UUID targetId);

    int denyAllIncomingFriendRequests(@NotNull UUID playerId);

    int denyAllOutgoingFriendRequests(@NotNull UUID playerId);

    @NotNull List<Friend> listFriends(@NotNull UUID playerId);

    @NotNull List<RequestedFriend> listPendingIncomingFriendRequests(@NotNull UUID playerId);

    @NotNull List<RequestedFriend> listPendingOutgoingFriendRequests(@NotNull UUID playerId);

    @NotNull CreateBlockResult block(@NotNull UUID blockerId, @NotNull UUID blockedId);

    @NotNull DeleteBlockResult unblock(@NotNull UUID blockerId, @NotNull UUID blockedId);

    @Nullable PlayerBlock isBlocked(@NotNull UUID blockerId, @NotNull UUID blockedId);

    @NotNull List<UUID> getBlockedList(@NotNull UUID blockerId);
}
