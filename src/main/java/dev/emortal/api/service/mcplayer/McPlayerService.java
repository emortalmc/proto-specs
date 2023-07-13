package dev.emortal.api.service.mcplayer;

import dev.emortal.api.grpc.mcplayer.McPlayerProto;
import dev.emortal.api.grpc.mcplayer.McPlayerProto.SearchPlayersByUsernameRequest.FilterMethod;
import dev.emortal.api.model.common.Pageable;
import dev.emortal.api.model.mcplayer.McPlayer;
import org.jetbrains.annotations.Blocking;
import org.jetbrains.annotations.NotNull;
import org.jetbrains.annotations.Nullable;

import java.util.Collection;
import java.util.List;
import java.util.UUID;

/**
 * A service interface for the mc player service.
 *
 * <p>
 * All methods within this class are blocking. It is up to the client to call these methods asynchronously.
 * <br>
 * The recommended way to do this is to use {@linkplain Thread#startVirtualThread virtual threads}.
 */
@Blocking
public interface McPlayerService {

    @Nullable McPlayer getPlayerById(@NotNull UUID id);

    @NotNull List<McPlayer> getPlayersById(@NotNull Collection<UUID> ids);

    @Nullable McPlayer getPlayerByUsername(@NotNull String username);

    @NotNull List<McPlayer> searchPlayersByUsername(@NotNull UUID requesterId, @NotNull String searchUsername, @NotNull Pageable pageable,
                                                    @Nullable FilterMethod filterMethod);

    @NotNull McPlayerProto.LoginSessionsResponse getLoginSessions(@NotNull UUID playerId, @NotNull Pageable pageable);
}
