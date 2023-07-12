package dev.emortal.api.service.relationship;

import dev.emortal.api.grpc.relationship.RelationshipProto.CreateBlockResponse.CreateBlockResult;
import dev.emortal.api.grpc.relationship.RelationshipProto.DenyFriendRequestResponse.DenyFriendRequestResult;
import dev.emortal.api.grpc.relationship.RelationshipProto.RemoveFriendResponse.RemoveFriendResult;
import dev.emortal.api.model.relationship.PlayerBlock;
import org.jetbrains.annotations.Blocking;
import org.jetbrains.annotations.NotNull;
import org.jetbrains.annotations.Nullable;

import java.util.List;
import java.util.UUID;

/**
 * A service interface for the relationship service.
 *
 * <p>
 * All methods within this class are blocking. It is up to the client to call these methods asynchronously.
 * <br>
 * The recommended way to do this is to use {@linkplain Thread#startVirtualThread virtual threads}.
 */
@Blocking
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
