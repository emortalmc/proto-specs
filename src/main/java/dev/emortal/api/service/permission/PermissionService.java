package dev.emortal.api.service.permission;

import dev.emortal.api.model.permission.Role;
import org.jetbrains.annotations.Blocking;
import org.jetbrains.annotations.NotNull;
import org.jetbrains.annotations.Nullable;

import java.util.List;
import java.util.UUID;

/**
 * A service interface for the permission service.
 *
 * <p>
 * All methods within this class are blocking. It is up to the client to call these methods asynchronously.
 * <br>
 * The recommended way to do this is to use {@linkplain Thread#startVirtualThread virtual threads}.
 */
@Blocking
public interface PermissionService {

    @NotNull List<Role> getAllRoles();

    @NotNull PlayerRoles getPlayerRoles(@NotNull UUID playerId);

    @NotNull CreateRoleResult createRole(@NotNull String id, int priority, @Nullable String displayName);

    @NotNull UpdateRoleResult updateRole(@NotNull RoleUpdate roleUpdate);

    @NotNull AddRoleToPlayerResult addRoleToPlayer(@NotNull UUID playerId, @NotNull String roleId);

    @NotNull RemoveRoleFromPlayerResult removeRoleFromPlayer(@NotNull UUID playerId, @NotNull String roleId);
}
