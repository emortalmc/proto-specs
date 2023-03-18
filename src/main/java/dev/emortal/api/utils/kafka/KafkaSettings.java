package dev.emortal.api.utils.kafka;

import lombok.Builder;
import lombok.Getter;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.net.InetAddress;
import java.net.UnknownHostException;

@Getter
@Builder
public class KafkaSettings {
    private static final Logger LOGGER = LoggerFactory.getLogger(KafkaSettings.class);

    // Required admin client
    private final String bootstrapServers;

    // Optional admin client
    @Builder.Default
    private String clientId = createClientId();
    @Builder.Default
    private String groupId = null;

    // Optional consumer
    @Builder.Default
    private boolean autoCommit = true;

    private static String createClientId() {
        try {
            return InetAddress.getLocalHost().getHostName();
        } catch (UnknownHostException e) {
            LOGGER.error("Failed to get hostname", e);
            return "unknown";
        }
    }
}
