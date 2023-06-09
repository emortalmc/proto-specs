package dev.emortal.api.utils;


import dev.emortal.api.grpc.accountconnmanager.AccountConnectionManagerGrpc;
import dev.emortal.api.grpc.badge.BadgeManagerGrpc;
import dev.emortal.api.grpc.gameplayerdata.GamePlayerDataServiceGrpc;
import dev.emortal.api.grpc.mcplayer.McPlayerGrpc;
import dev.emortal.api.grpc.messagehandler.MessageHandlerGrpc;
import dev.emortal.api.grpc.party.PartyServiceGrpc;
import dev.emortal.api.grpc.party.PartySettingsServiceGrpc;
import dev.emortal.api.grpc.permission.PermissionServiceGrpc;
import dev.emortal.api.grpc.playertracker.PlayerTrackerGrpc;
import dev.emortal.api.grpc.relationship.RelationshipGrpc;
import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;
import lombok.Getter;
import org.jetbrains.annotations.NotNull;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.io.IOException;
import java.net.ConnectException;
import java.net.Socket;
import java.util.Map;
import java.util.Optional;

// We directly return the optionals here, so it's functionally identical
@SuppressWarnings("OptionalUsedAsFieldOrParameterType")
public class GrpcStubCollection {
    private static final Logger LOGGER = LoggerFactory.getLogger(GrpcStubCollection.class);
    private static final boolean DEVELOPMENT = System.getenv("KUBERNETES_SERVICE_HOST") == null;

    private static final Map<String, Integer> LOCAL_SERVICE_PORTS = Map.of(
            "permission", 10001,
            "relationship-manager", 10002,
            "message-handler", 10003,
            "mc-player", 10004,
            "badge-manager", 10004,
            "player-tracker", 10005,
            "party-manager", 10006,
            // Kurushimi is 10007
            "account-connection-manager", 10008,
            "game-player-data", 10009
    );

    @Getter
    private static final @NotNull Optional<PermissionServiceGrpc.PermissionServiceFutureStub> permissionService;
    @Getter
    private static final @NotNull Optional<RelationshipGrpc.RelationshipFutureStub> relationshipService;
    @Getter
    private static final @NotNull Optional<MessageHandlerGrpc.MessageHandlerFutureStub> messageHandlerService;
    @Getter
    private static final @NotNull Optional<McPlayerGrpc.McPlayerFutureStub> playerService;
    @Getter
    private static final @NotNull Optional<BadgeManagerGrpc.BadgeManagerFutureStub> badgeManagerService;
    @Getter
    private static final @NotNull Optional<PlayerTrackerGrpc.PlayerTrackerFutureStub> playerTrackerService;
    @Getter
    private static final @NotNull Optional<PartyServiceGrpc.PartyServiceFutureStub> partyService;
    @Getter
    private static final @NotNull Optional<PartySettingsServiceGrpc.PartySettingsServiceFutureStub> partySettingsService;
    @Getter
    private static final @NotNull Optional<AccountConnectionManagerGrpc.AccountConnectionManagerFutureStub> accountConnectionManagerService;
    @Getter
    private static final @NotNull Optional<GamePlayerDataServiceGrpc.GamePlayerDataServiceFutureStub> gamePlayerDataService;

    static {
        permissionService = createChannel("permission").map(PermissionServiceGrpc::newFutureStub);
        relationshipService = createChannel("relationship-manager").map(RelationshipGrpc::newFutureStub);
        messageHandlerService = createChannel("message-handler").map(MessageHandlerGrpc::newFutureStub);
        playerService = createChannel("mc-player").map(McPlayerGrpc::newFutureStub);
        badgeManagerService = createChannel("badge-manager").map(BadgeManagerGrpc::newFutureStub);
        playerTrackerService = createChannel("player-tracker").map(PlayerTrackerGrpc::newFutureStub);
        partyService = createChannel("party-manager").map(PartyServiceGrpc::newFutureStub);
        partySettingsService = createChannel("party-manager").map(PartySettingsServiceGrpc::newFutureStub);
        accountConnectionManagerService = createChannel("account-connection-manager").map(AccountConnectionManagerGrpc::newFutureStub);
        gamePlayerDataService = createChannel("game-player-data").map(GamePlayerDataServiceGrpc::newFutureStub);
    }

    /**
     * @param name Service name in the format "svc-name"
     * @return Optional of a ManagedChannel, empty if service is not enabled.
     */
    private static Optional<ManagedChannel> createChannel(String name) {
        if (!DEVELOPMENT) {
            return Optional.of(ManagedChannelBuilder.forAddress(name, 9010)
                    .defaultLoadBalancingPolicy("round_robin")
                    .usePlaintext()
                    .build());
        }
        String envPort = System.getenv(name + "-port"); // only used for dev

        if (envPort == null) {
            int defaultPort = LOCAL_SERVICE_PORTS.get(name);
            if (!isPortUsed(defaultPort)) {
                LOGGER.warn("Service {} is not enabled (port: {}), skipping", name, defaultPort);
                return Optional.empty();
            }

            LOGGER.info("Automatically using port {} for service {}", defaultPort, name);
            return Optional.of(ManagedChannelBuilder.forAddress("localhost", defaultPort)
                    .usePlaintext()
                    .build());
        } else {
            int port = Integer.parseInt(envPort);
            return Optional.of(ManagedChannelBuilder.forAddress("localhost", port)
                    .usePlaintext()
                    .build());
        }
    }

    /**
     * @param port Port to check
     * @return True if the port is used, false if not.
     */
    private static boolean isPortUsed(int port) {
        try (Socket socket = new Socket("localhost", port)) {
            socket.setSoTimeout(10);
            return true;
        } catch (ConnectException e) {
            return false;
        } catch (IOException e) {
            LOGGER.error("Error while checking if port is used", e);
            return false;
        }
    }
}
