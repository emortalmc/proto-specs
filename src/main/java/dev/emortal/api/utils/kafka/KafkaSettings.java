package dev.emortal.api.utils.kafka;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.net.InetAddress;
import java.net.UnknownHostException;
import java.util.HashMap;
import java.util.Map;

public class KafkaSettings {
    private static final Logger LOGGER = LoggerFactory.getLogger(KafkaSettings.class);

    private final Map<String, Object> settings = new HashMap<>();
    private boolean autoCommit = true;

    public KafkaSettings() {
        this.settings.put("client.id", this.createClientId());
    }

    public KafkaSettings setClientId(String clientId) {
        this.settings.put("client.id", clientId);
        return this;
    }

    public KafkaSettings setBootstrapServers(String bootstrapServers) {
        this.settings.put("bootstrap.servers", bootstrapServers);
        return this;
    }

    public KafkaSettings setGroupId(String groupId) {
        this.settings.put("group.id", groupId);
        return this;
    }

    public KafkaSettings setAutoCommit(boolean autoCommit) {
        this.autoCommit = autoCommit;
        return this;
    }

    public boolean isAutoCommit() {
        return this.autoCommit;
    }

    public Map<String, Object> getSettings() {
        if (this.settings.get("bootstrap.servers") == null)
            throw new IllegalStateException("Bootstrap servers not set");

        return this.settings;
    }

    private String createClientId() {
        try {
            return InetAddress.getLocalHost().getHostName();
        } catch (UnknownHostException e) {
            LOGGER.error("Failed to get hostname", e);
            return "unknown";
        }
    }
}
