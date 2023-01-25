package dev.emortal.api.utils;


import dev.emortal.api.grpc.friend.FriendGrpc;
import dev.emortal.api.grpc.mcplayer.McPlayerGrpc;
import dev.emortal.api.grpc.permission.PermissionServiceGrpc;
import dev.emortal.api.grpc.playertracker.PlayerTrackerGrpc;
import dev.emortal.api.grpc.privatemessage.PrivateMessageGrpc;
import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;
import lombok.Getter;
import org.jetbrains.annotations.NotNull;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.util.Optional;

// We directly return the optionals here, so it's functionally identical
@SuppressWarnings("OptionalUsedAsFieldOrParameterType")
public class GrpcStubCollection {
    private static final Logger LOGGER = LoggerFactory.getLogger(GrpcStubCollection.class);
    private static final boolean DEVELOPMENT = System.getenv("HOSTNAME") == null;

    @Getter
    private static final @NotNull Optional<FriendGrpc.FriendFutureStub> friendService;
    @Getter
    private static final @NotNull Optional<McPlayerGrpc.McPlayerFutureStub> playerService;
    @Getter
    private static final @NotNull Optional<PermissionServiceGrpc.PermissionServiceFutureStub> permissionService;
    @Getter
    private static final @NotNull Optional<PlayerTrackerGrpc.PlayerTrackerFutureStub> playerTrackerService;
    @Getter
    private static final @NotNull Optional<PrivateMessageGrpc.PrivateMessageFutureStub> privateMessageService;

    static {
        friendService = createChannel("friend-manager").map(FriendGrpc::newFutureStub);
        playerService = createChannel("mc-player").map(McPlayerGrpc::newFutureStub);
        permissionService = createChannel("permission").map(PermissionServiceGrpc::newFutureStub);
        playerTrackerService = createChannel("player-tracker").map(PlayerTrackerGrpc::newFutureStub);
        privateMessageService = createChannel("private-message").map(PrivateMessageGrpc::newFutureStub);
    }

    /**
     * @param name Service name in the format "svc-name"
     * @return Optional of a ManagedChannel, empty if service is not enabled.
     */
    private static Optional<ManagedChannel> createChannel(String name) {
        if (!DEVELOPMENT) {
            return Optional.of(ManagedChannelBuilder.forAddress(name, 9090)
                    .defaultLoadBalancingPolicy("round_robin")
                    .usePlaintext()
                    .build());
        }
        String envPort = System.getenv(name + "-port"); // only used for dev

        if (envPort == null) {
            LOGGER.warn("Service {} is not enabled, skipping", name);
            return Optional.empty();
        } else {
            int port = Integer.parseInt(envPort);
            return Optional.of(ManagedChannelBuilder.forAddress("localhost", port)
                    .usePlaintext()
                    .build());
        }
    }
}