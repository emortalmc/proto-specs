package dev.emortal.api.service.permission;

import dev.emortal.api.model.permission.PermissionNode;
import org.jetbrains.annotations.NotNull;
import org.jetbrains.annotations.Nullable;

import java.util.ArrayList;
import java.util.List;

public final class RoleUpdate {

    public static @NotNull Builder builder(@NotNull String id) {
        return new Builder(id);
    }

    private final String id;
    private final int priority;
    private final @Nullable String displayName;
    private final List<PermissionNode> setPermissions;
    private final List<String> unsetPermissions;

    RoleUpdate(@NotNull Builder builder) {
        this.id = builder.id;
        this.priority = builder.priority;
        this.displayName = builder.displayName;
        this.setPermissions = builder.setPermissions;
        this.unsetPermissions = builder.unsetPermissions;
    }

    @NotNull String id() {
        return this.id;
    }

    int priority() {
        return this.priority;
    }

    @Nullable String displayName() {
        return this.displayName;
    }

    @NotNull List<PermissionNode> setPermissions() {
        return this.setPermissions;
    }

    @NotNull List<String> unsetPermissions() {
        return this.unsetPermissions;
    }

    public static final class Builder {

        private final String id;
        private int priority = -1;
        private @Nullable String displayName;
        private final List<PermissionNode> setPermissions = new ArrayList<>();
        private final List<String> unsetPermissions = new ArrayList<>();

        private Builder(@NotNull String id) {
            this.id = id;
        }

        public @NotNull Builder priority(int priority) {
            if (priority < 0) throw new IllegalArgumentException("Priority cannot be less than 0");
            this.priority = priority;
            return this;
        }

        public @NotNull Builder displayName(@Nullable String displayName) {
            this.displayName = displayName;
            return this;
        }

        public @NotNull Builder setPermission(@NotNull PermissionNode permission) {
            this.setPermissions.add(permission);
            return this;
        }

        public @NotNull Builder unsetPermission(@NotNull String permission) {
            this.unsetPermissions.add(permission);
            return this;
        }

        public @NotNull RoleUpdate build() {
            return new RoleUpdate(this);
        }
    }
}
