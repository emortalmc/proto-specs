package dev.emortal.api.service.playertracker;

import dev.emortal.api.model.mcplayer.CurrentServer;
import dev.emortal.api.model.mcplayer.OnlinePlayer;
import org.jetbrains.annotations.Blocking;
import org.jetbrains.annotations.NotNull;
import org.jetbrains.annotations.Nullable;

import java.util.*;

/**
 * A service interface for the player tracker.
 *
 * <p>
 * All methods within this class are blocking. It is up to the client to call these methods asynchronously.
 * <br>
 * The recommended way to do this is to use {@linkplain Thread#startVirtualThread virtual threads}.
 */
@Blocking
public interface PlayerTrackerService {

    @NotNull Map<UUID, CurrentServer> getServersByPlayers(@NotNull Collection<UUID> playerIds);

    @Nullable CurrentServer getServer(@NotNull UUID playerId);

    @NotNull List<OnlinePlayer> getServerPlayers(@NotNull String serverId);

    long getPlayerCount(@NotNull String serverId, @NotNull Collection<String> fleetNames);

    long getPlayerCount(@NotNull String serverId);

    long getGlobalPlayerCount();

    @NotNull Map<String, Long> getFleetPlayerCounts(@NotNull Collection<String> fleetNames);

    @NotNull List<OnlinePlayer> getGlobalPlayersSummary(@Nullable String serverId, @Nullable Set<String> fleetNames);

    default @NotNull List<OnlinePlayer> getGlobalPlayersSummaryByServer(@NotNull String serverId) {
        return this.getGlobalPlayersSummary(serverId, null);
    }

    default @NotNull List<OnlinePlayer> getGlobalPlayersSummaryByFleet(@NotNull Set<String> fleetNames) {
        return this.getGlobalPlayersSummary(null, fleetNames);
    }

    default @NotNull List<OnlinePlayer> getGlobalPlayersSummary() {
        return this.getGlobalPlayersSummary(null, null);
    }
}
