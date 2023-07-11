package dev.emortal.api.service.permission;

import org.jetbrains.annotations.NotNull;
import org.jetbrains.annotations.Nullable;

import java.util.List;

public record PlayerRoles(@NotNull List<String> roleIds, @Nullable String activeDisplayNameRoleId) {
}
