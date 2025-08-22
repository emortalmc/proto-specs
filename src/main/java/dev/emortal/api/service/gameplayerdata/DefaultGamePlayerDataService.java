package dev.emortal.api.service.gameplayerdata;

import com.google.protobuf.Any;
import com.google.protobuf.InvalidProtocolBufferException;
import com.google.protobuf.Message;
import dev.emortal.api.grpc.gameplayerdata.GamePlayerDataProto;
import dev.emortal.api.grpc.gameplayerdata.GamePlayerDataServiceGrpc;
import dev.emortal.api.grpc.gameplayerdata.GamePlayerDataProto.SetMinesweeperProfileResponse;
import dev.emortal.api.grpc.matchmaker.MatchmakerGrpc;
import dev.emortal.api.model.gamedata.GameDataGameMode;
import dev.emortal.api.model.gamedata.V1MinesweeperProfile;
import io.grpc.Channel;
import io.grpc.StatusRuntimeException;
import org.jetbrains.annotations.ApiStatus;
import org.jetbrains.annotations.NotNull;
import org.jetbrains.annotations.Nullable;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.util.HashMap;
import java.util.Map;
import java.util.UUID;
import java.util.stream.StreamSupport;

public class DefaultGamePlayerDataService implements GamePlayerDataService {
    private static final Logger LOGGER = LoggerFactory.getLogger(DefaultGamePlayerDataService.class);

    private final @NotNull GamePlayerDataServiceGrpc.GamePlayerDataServiceBlockingStub grpc;

    @ApiStatus.Internal
    public DefaultGamePlayerDataService(@NotNull Channel channel) {
        this.grpc = GamePlayerDataServiceGrpc.newBlockingStub(channel);
    }

    @Override
    public <T extends Message> @Nullable T getGameData(@NotNull GameDataGameMode gameMode, @NotNull Class<T> clazz, @NotNull UUID playerId) {
        var request = GamePlayerDataProto.GetGamePlayerDataRequest.newBuilder().setGameMode(gameMode)
                .setPlayerId(playerId.toString()).build();

        try {
            GamePlayerDataProto.GetGamePlayerDataResponse response = this.grpc.getGamePlayerData(request);

            return response.getData().unpack(clazz);
        } catch (StatusRuntimeException exception) {
            if (exception.getStatus().getCode() == io.grpc.Status.Code.NOT_FOUND) return null;
            throw exception;
        } catch (InvalidProtocolBufferException e) {
            LOGGER.error("Error unpacking game data for type {}: ", clazz.getSimpleName(), e);
            return null;
        }
    }

    @Override
    public @NotNull <T extends Message> Map<UUID, @Nullable T> getGameData(@NotNull GameDataGameMode gameMode, @NotNull Class<T> clazz, @NotNull Iterable<@NotNull UUID> playerIds) {
        var request = GamePlayerDataProto.GetMultipleGamePlayerDataRequest.newBuilder().setGameMode(gameMode)
                .addAllPlayerIds(StreamSupport.stream(playerIds.spliterator(), false).map(UUID::toString).toList()).build();

        GamePlayerDataProto.GetMultipleGamePlayerDataResponse response = this.grpc.getMultipleGamePlayerData(request);
        Map<String, Any> responseMap = response.getDataMap();

        Map<UUID, @Nullable T> returnMap = new HashMap<>();
        for (UUID playerId : playerIds) {
            returnMap.put(playerId, null);
        }

        for (Map.Entry<String, Any> entry : responseMap.entrySet()) {
            try {
                returnMap.put(UUID.fromString(entry.getKey()), entry.getValue().unpack(clazz));
            } catch (InvalidProtocolBufferException e) {
                LOGGER.error("Error unpacking game data for type {}: ", clazz.getSimpleName(), e);
            }
        }

        return returnMap;
    }

    @Override
    public String setMinesweeperProfile(@NotNull V1MinesweeperProfile profile) {
        var request = GamePlayerDataProto.SetMinesweeperProfileRequest.newBuilder().setProfile(profile).build();

        GamePlayerDataProto.SetMinesweeperProfileResponse response = this.grpc.setMinesweeperProfile(request);
        return response.getProfileId();

    }
    @Override
    public V1MinesweeperProfile getMinesweeperProfile(String profileId) {
        var request = GamePlayerDataProto.GetMinesweeperProfileRequest.newBuilder().setProfileId(profileId).build();

        GamePlayerDataProto.GetMinesweeperProfileResponse response = this.grpc.getMinesweeperProfile(request);
        return response.getProfile();
    }

    @Override
    public V1MinesweeperProfile[] listMinesweeperProfile(String ownerId) {
        var request = GamePlayerDataProto.ListMinesweeperProfilesRequest.newBuilder().setOwnerId(ownerId).build();

        GamePlayerDataProto.ListMinesweeperProfilesResponse response = this.grpc.listMinesweeperProfiles(request);
        return response.getProfileList().toArray(new V1MinesweeperProfile[response.getProfileCount()]);
    }
}

