package dev.emortal.api.service.permission;

import dev.emortal.api.model.permission.Role;
import org.jetbrains.annotations.NotNull;

public sealed interface CreateRoleResult {

    record Success(@NotNull Role role) implements CreateRoleResult {
    }

    enum Error implements CreateRoleResult {

        ROLE_ALREADY_EXISTS
    }
}
