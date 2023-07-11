package dev.emortal.api.service.playertracker;

import dev.emortal.api.grpc.mcplayer.PlayerTrackerGrpc;
import dev.emortal.api.grpc.mcplayer.McPlayerProto;
import dev.emortal.api.model.mcplayer.CurrentServer;
import dev.emortal.api.model.mcplayer.OnlinePlayer;
import io.grpc.ManagedChannel;
import org.jetbrains.annotations.ApiStatus;
import org.jetbrains.annotations.NotNull;
import org.jetbrains.annotations.Nullable;

import java.util.Collection;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.UUID;

@ApiStatus.Internal
public final class DefaultPlayerTrackerService implements PlayerTrackerService {

    private final PlayerTrackerGrpc.PlayerTrackerBlockingStub grpc;

    @ApiStatus.Internal
    public DefaultPlayerTrackerService(@NotNull ManagedChannel channel) {
        this.grpc = PlayerTrackerGrpc.newBlockingStub(channel);
    }

    @Override
    public @NotNull Map<UUID, CurrentServer> getServersByPlayers(@NotNull Collection<UUID> playerIds) {
        var requestBuilder = McPlayerProto.GetPlayerServersRequest.newBuilder();
        for (UUID id : playerIds) {
            requestBuilder.addPlayerIds(id.toString());
        }
        var request = requestBuilder.build();

        McPlayerProto.GetPlayerServersResponse response = this.grpc.getPlayerServers(request);
        Map<UUID, CurrentServer> result = new HashMap<>();
        for (var entry : response.getPlayerServersMap().entrySet()) {
            result.put(UUID.fromString(entry.getKey()), entry.getValue());
        }
        return result;
    }

    @Override
    public @Nullable CurrentServer getServer(@NotNull UUID playerId) {
        var request = McPlayerProto.GetPlayerServersRequest.newBuilder().addPlayerIds(playerId.toString()).build();

        McPlayerProto.GetPlayerServersResponse response = this.grpc.getPlayerServers(request);
        return response.getPlayerServersMap().get(playerId.toString());
    }

    @Override
    public @NotNull List<OnlinePlayer> getServerPlayers(@NotNull String serverId) {
        var request = McPlayerProto.GetServerPlayersRequest.newBuilder().setServerId(serverId).build();

        McPlayerProto.GetServerPlayersResponse response = this.grpc.getServerPlayers(request);
        return response.getOnlinePlayersList();
    }

    @Override
    public long getPlayerCount(@NotNull String serverId, @NotNull Collection<String> fleetNames) {
        return getPlayerCount0(serverId, fleetNames);
    }

    @Override
    public long getPlayerCount(@NotNull String serverId) {
        return getPlayerCount0(serverId, null);
    }

    @Override
    public long getGlobalPlayerCount() {
        return getPlayerCount0(null, null);
    }

    private long getPlayerCount0(@Nullable String serverId, @Nullable Collection<String> fleetNames) {
        var requestBuilder = McPlayerProto.GetPlayerCountRequest.newBuilder();
        if (serverId != null) requestBuilder.setServerId(serverId);
        if (fleetNames != null) requestBuilder.addAllFleetNames(fleetNames);
        var request = requestBuilder.build();

        McPlayerProto.GetPlayerCountResponse response = this.grpc.getPlayerCount(request);
        return response.getCount();
    }

    @Override
    public @NotNull Map<String, Long> getFleetPlayerCounts(@NotNull Collection<String> fleetNames) {
        var request = McPlayerProto.GetFleetsPlayerCountRequest.newBuilder().addAllFleetNames(fleetNames).build();

        McPlayerProto.GetFleetsPlayerCountResponse response = this.grpc.getFleetPlayerCounts(request);
        return response.getFleetPlayerCountsMap();
    }
}
