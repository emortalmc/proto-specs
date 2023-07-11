package dev.emortal.api.service.permission;

import dev.emortal.api.model.permission.Role;
import org.jetbrains.annotations.NotNull;
import org.jetbrains.annotations.Nullable;

import java.util.List;
import java.util.UUID;

public interface PermissionService {

    @NotNull List<Role> getAllRoles();

    @NotNull PlayerRoles getPlayerRoles(@NotNull UUID playerId);

    @NotNull CreateRoleResult createRole(@NotNull String id, int priority, @Nullable String displayName);

    @NotNull UpdateRoleResult updateRole(@NotNull RoleUpdate roleUpdate);

    @NotNull AddRoleToPlayerResult addRoleToPlayer(@NotNull UUID playerId, @NotNull String roleId);

    @NotNull RemoveRoleFromPlayerResult removeRoleFromPlayer(@NotNull UUID playerId, @NotNull String roleId);
}
