package dev.emortal.api.utils.kafka;

import org.jetbrains.annotations.NotNull;
import org.jetbrains.annotations.Nullable;

import java.net.InetAddress;
import java.net.UnknownHostException;
import java.util.Arrays;
import java.util.HashMap;
import java.util.HashSet;
import java.util.Map;
import java.util.Set;

record KafkaSettings(@NotNull String clientId, @NotNull Set<String> bootstrapServers, @Nullable String groupId,
                     boolean autoCommit) {

    private static final String CLIENT_ID_KEY = "client.id";
    private static final String BOOTSTRAP_SERVERS_KEY = "bootstrap.servers";
    private static final String GROUP_ID_KEY = "group.id";

    public @NotNull Map<@NotNull String, @NotNull Object> getSettings() {
        Map<String, Object> settings = new HashMap<>();
        settings.put(CLIENT_ID_KEY, this.clientId);
        settings.put(BOOTSTRAP_SERVERS_KEY, String.join(",", this.bootstrapServers));

        if (this.groupId != null) settings.put(GROUP_ID_KEY, this.groupId);

        return settings;
    }

    public static class Builder {
        private @NotNull String clientId;
        private Set<String> bootstrapServers;
        private String groupId;
        private boolean autoCommit;

        public Builder() {
            this.clientId = this.createClientId();
        }

        public @NotNull Builder clientId(@NotNull String clientId) {
            this.clientId = clientId;
            return this;
        }

        public @NotNull Builder bootstrapServers(@NotNull Set<String> bootstrapServers) {
            this.bootstrapServers = bootstrapServers;
            return this;
        }

        public @NotNull Builder bootstrapServers(@NotNull String... bootstrapServers) {
            if (this.bootstrapServers == null) this.bootstrapServers = new HashSet<>();
            this.bootstrapServers.addAll(Arrays.asList(bootstrapServers));
            return this;
        }

        public @NotNull Builder groupId(@NotNull String groupId) {
            this.groupId = groupId;
            return this;
        }

        public @NotNull Builder autoCommit(boolean autoCommit) {
            this.autoCommit = autoCommit;
            return this;
        }

        public @NotNull KafkaSettings build() {
            if (this.bootstrapServers == null || this.bootstrapServers.isEmpty())
                throw new IllegalStateException("Bootstrap servers not set");

            // Only set the group id if it's null and auto commit is enabled
            if (this.groupId == null && this.autoCommit)
                this.groupId = this.clientId;

            return new KafkaSettings(this.clientId, this.bootstrapServers, this.groupId, this.autoCommit);
        }

        private String createClientId() {
            try {
                return InetAddress.getLocalHost().getHostName();
            } catch (UnknownHostException e) {
                throw new RuntimeException(e);
            }
        }
    }
}
