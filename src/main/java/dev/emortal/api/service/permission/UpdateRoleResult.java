package dev.emortal.api.service.permission;

import dev.emortal.api.model.permission.Role;
import org.jetbrains.annotations.NotNull;

public sealed interface UpdateRoleResult {

    record Success(@NotNull Role role) implements UpdateRoleResult {
    }

    enum Error implements UpdateRoleResult {

        ROLE_NOT_FOUND
    }
}
