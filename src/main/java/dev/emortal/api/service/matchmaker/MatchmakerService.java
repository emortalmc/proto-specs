package dev.emortal.api.service.matchmaker;

import dev.emortal.api.grpc.matchmaker.MatchmakerProto;
import org.jetbrains.annotations.Blocking;
import org.jetbrains.annotations.NotNull;
import org.jetbrains.annotations.Nullable;

import java.util.Collection;
import java.util.UUID;

/**
 * A service interface for the matchmaker service.
 *
 * <p>
 * All methods within this class are blocking. It is up to the client to call these methods asynchronously.
 * <br>
 * The recommended way to do this is to use {@linkplain Thread#startVirtualThread virtual threads}.
 */
@Blocking
public interface MatchmakerService {

    @Nullable MatchmakerProto.GetPlayerQueueInfoResponse getPlayerQueueInfo(@NotNull UUID playerId);

    @NotNull QueuePlayerResult queuePlayer(@NotNull String gameModeId, @NotNull UUID playerId, @NotNull QueueOptions options);

    default @NotNull QueuePlayerResult queuePlayer(@NotNull String gameModeId, @NotNull UUID playerId) {
        return this.queuePlayer(gameModeId, playerId, QueueOptions.DEFAULT);
    }

    @NotNull DequeuePlayerResult dequeuePlayer(@NotNull UUID playerId);

    void sendPlayerToLobby(@NotNull UUID playerId, boolean sendParty);

    void sendPlayersToLobby(@NotNull Collection<UUID> playerIds, boolean sendParties);

    void loginQueue(@NotNull UUID playerId, boolean proxy);

    @NotNull ChangeMapVoteResult changeMapVote(@NotNull UUID playerId, @NotNull String newMapId);
}
