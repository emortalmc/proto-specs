package dev.emortal.api.service.gameplayerdata;

import com.google.protobuf.Message;
import dev.emortal.api.model.gamedata.GameDataGameMode;
import dev.emortal.api.model.gamedata.V1MinesweeperProfile;

import org.jetbrains.annotations.NotNull;
import org.jetbrains.annotations.Nullable;

import java.util.Map;
import java.util.UUID;

public interface GamePlayerDataService {

    <T extends Message> @Nullable T getGameData(@NotNull GameDataGameMode gameMode, @NotNull Class<T> clazz, @NotNull UUID playerId);

    <T extends Message> @NotNull Map<UUID, @Nullable T> getGameData(@NotNull GameDataGameMode gameMode, @NotNull Class<T> clazz, @NotNull Iterable<@NotNull UUID> playerIds);

    String setMinesweeperProfile(@NotNull V1MinesweeperProfile profile);
    V1MinesweeperProfile getMinesweeperProfile(@Nullable String profileId);
    V1MinesweeperProfile[] listMinesweeperProfile(@Nullable String ownerId);
}
