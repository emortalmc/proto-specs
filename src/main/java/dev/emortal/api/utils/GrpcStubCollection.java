package dev.emortal.api.utils;

import dev.emortal.api.grpc.accountconnmanager.AccountConnectionManagerGrpc;
import dev.emortal.api.grpc.gameplayerdata.GamePlayerDataServiceGrpc;
import dev.emortal.api.grpc.party.PartySettingsServiceGrpc;
import dev.emortal.api.service.badges.BadgeService;
import dev.emortal.api.service.badges.DefaultBadgeService;
import dev.emortal.api.service.gameplayerdata.DefaultGamePlayerDataService;
import dev.emortal.api.service.gameplayerdata.GamePlayerDataService;
import dev.emortal.api.service.matchmaker.DefaultMatchmakerService;
import dev.emortal.api.service.matchmaker.MatchmakerService;
import dev.emortal.api.service.mcplayer.DefaultMcPlayerService;
import dev.emortal.api.service.mcplayer.McPlayerService;
import dev.emortal.api.service.messagehandler.DefaultMessageService;
import dev.emortal.api.service.messagehandler.MessageService;
import dev.emortal.api.service.party.DefaultPartyService;
import dev.emortal.api.service.party.PartyService;
import dev.emortal.api.service.permission.DefaultPermissionService;
import dev.emortal.api.service.permission.PermissionService;
import dev.emortal.api.service.playertracker.DefaultPlayerTrackerService;
import dev.emortal.api.service.playertracker.PlayerTrackerService;
import dev.emortal.api.service.relationship.DefaultRelationshipService;
import dev.emortal.api.service.relationship.RelationshipService;
import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;
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
public final class GrpcStubCollection {
    private static final Logger LOGGER = LoggerFactory.getLogger(GrpcStubCollection.class);
    private static final boolean DEVELOPMENT = System.getenv("KUBERNETES_SERVICE_HOST") == null;

    private static final Map<String, Integer> LOCAL_SERVICE_PORTS = Map.of(
            "permission", 10001,
            "relationship-manager", 10002,
            "message-handler", 10003,
            "mc-player", 10004,
            "player-tracker", 10004,
            "badge-manager", 10004,
            "party-manager", 10006,
            "matchmaker", 10007,
            "account-connection-manager", 10008,
            "game-player-data", 10009
    );

    private static final @NotNull Optional<PermissionService> permissionService;
    private static final @NotNull Optional<RelationshipService> relationshipService;
    private static final @NotNull Optional<MessageService> messageHandlerService;
    private static final @NotNull Optional<McPlayerService> playerService;
    private static final @NotNull Optional<PlayerTrackerService> playerTrackerService;
    private static final @NotNull Optional<BadgeService> badgeManagerService;
    private static final @NotNull Optional<PartyService> partyService;
    private static final @NotNull Optional<PartySettingsServiceGrpc.PartySettingsServiceBlockingStub> partySettingsService;
    private static final @NotNull Optional<MatchmakerService> matchmakerService;
    private static final @NotNull Optional<AccountConnectionManagerGrpc.AccountConnectionManagerBlockingStub> accountConnectionManagerService;
    private static final @NotNull Optional<GamePlayerDataService> gamePlayerDataService;

    static {
        permissionService = createChannel("permission").map(DefaultPermissionService::new);
        relationshipService = createChannel("relationship-manager").map(DefaultRelationshipService::new);
        messageHandlerService = createChannel("message-handler").map(DefaultMessageService::new);
        playerService = createChannel("mc-player").map(DefaultMcPlayerService::new);
        playerTrackerService = createChannel("player-tracker").map(DefaultPlayerTrackerService::new);
        badgeManagerService = createChannel("badge-manager").map(DefaultBadgeService::new);
        partyService = createChannel("party-manager").map(DefaultPartyService::new);
        partySettingsService = createChannel("party-manager").map(PartySettingsServiceGrpc::newBlockingStub);
        matchmakerService = createChannel("matchmaker").map(DefaultMatchmakerService::new);
        accountConnectionManagerService = createChannel("account-connection-manager").map(AccountConnectionManagerGrpc::newBlockingStub);
        gamePlayerDataService = createChannel("game-player-data").map(DefaultGamePlayerDataService::new);

    }

    public static @NotNull Optional<PermissionService> getPermissionService() {
        return permissionService;
    }

    public static @NotNull Optional<RelationshipService> getRelationshipService() {
        return relationshipService;
    }

    public static @NotNull Optional<MessageService> getMessageHandlerService() {
        return messageHandlerService;
    }

    public static @NotNull Optional<McPlayerService> getPlayerService() {
        return playerService;
    }

    public static @NotNull Optional<PlayerTrackerService> getPlayerTrackerService() {
        return playerTrackerService;
    }

    public static @NotNull Optional<BadgeService> getBadgeManagerService() {
        return badgeManagerService;
    }

    public static @NotNull Optional<PartyService> getPartyService() {
        return partyService;
    }

    public static @NotNull Optional<PartySettingsServiceGrpc.PartySettingsServiceBlockingStub> getPartySettingsService() {
        return partySettingsService;
    }

    public static @NotNull Optional<MatchmakerService> getMatchmakerService() {
        return matchmakerService;
    }

    public static @NotNull Optional<AccountConnectionManagerGrpc.AccountConnectionManagerBlockingStub> getAccountConnectionManagerService() {
        return accountConnectionManagerService;
    }

    public static @NotNull Optional<GamePlayerDataService> getGamePlayerDataService() {
        return gamePlayerDataService;
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
        } catch (ConnectException exception) {
            return false;
        } catch (IOException exception) {
            LOGGER.error("Error while checking if port is used", exception);
            return false;
        }
    }

    private GrpcStubCollection() {
    }
}
