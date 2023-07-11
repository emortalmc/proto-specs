package dev.emortal.api.service.playertracker;

import dev.emortal.api.model.mcplayer.CurrentServer;
import dev.emortal.api.model.mcplayer.OnlinePlayer;
import org.jetbrains.annotations.NotNull;
import org.jetbrains.annotations.Nullable;

import java.util.Collection;
import java.util.List;
import java.util.Map;
import java.util.UUID;

public interface PlayerTrackerService {

    @NotNull Map<UUID, CurrentServer> getServersByPlayers(@NotNull Collection<UUID> playerIds);

    @Nullable CurrentServer getServer(@NotNull UUID playerId);

    @NotNull List<OnlinePlayer> getServerPlayers(@NotNull String serverId);

    long getPlayerCount(@NotNull String serverId, @NotNull Collection<String> fleetNames);

    long getPlayerCount(@NotNull String serverId);

    long getGlobalPlayerCount();

    @NotNull Map<String, Long> getFleetPlayerCounts(@NotNull Collection<String> fleetNames);
}
