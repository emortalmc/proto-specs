package dev.emortal.api.service.mcplayer;

import dev.emortal.api.grpc.mcplayer.McPlayerProto.SearchPlayersByUsernameRequest.FilterMethod;
import dev.emortal.api.model.common.Pageable;
import dev.emortal.api.model.mcplayer.McPlayer;
import org.jetbrains.annotations.NotNull;
import org.jetbrains.annotations.Nullable;

import java.util.Collection;
import java.util.List;
import java.util.UUID;

public interface McPlayerService {

    @Nullable McPlayer getPlayerById(@NotNull UUID id);

    @NotNull List<McPlayer> getPlayersById(@NotNull Collection<UUID> ids);

    @Nullable McPlayer getPlayerByUsername(@NotNull String username);

    @NotNull List<McPlayer> searchPlayersByUsername(@NotNull UUID requesterId, @NotNull String searchUsername, @NotNull Pageable pageable,
                                                    @Nullable FilterMethod filterMethod);

    @NotNull LoginSessions getLoginSessions(@NotNull UUID playerId, @NotNull Pageable pageable);
}
