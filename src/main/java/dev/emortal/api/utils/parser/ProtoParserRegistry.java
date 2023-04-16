package dev.emortal.api.utils.parser;

import com.google.protobuf.InvalidProtocolBufferException;
import com.google.protobuf.Message;
import dev.emortal.api.message.accountconnmanager.AccountConnectedMessage;
import dev.emortal.api.message.accountconnmanager.AccountConnectionRemovedMessage;
import dev.emortal.api.message.badge.PlayerActiveBadgeChangedMessage;
import dev.emortal.api.message.badge.PlayerBadgeAddedMessage;
import dev.emortal.api.message.badge.PlayerBadgeRemovedMessage;
import dev.emortal.api.message.common.PlayerConnectMessage;
import dev.emortal.api.message.common.PlayerDisconnectMessage;
import dev.emortal.api.message.common.PlayerSwitchServerMessage;
import dev.emortal.api.message.messagehandler.ChatMessageCreatedMessage;
import dev.emortal.api.message.messagehandler.PrivateMessageCreatedMessage;
import dev.emortal.api.message.party.PartyCreatedMessage;
import dev.emortal.api.message.party.PartyDeletedMessage;
import dev.emortal.api.message.party.PartyEmptiedMessage;
import dev.emortal.api.message.party.PartyInviteCreatedMessage;
import dev.emortal.api.message.party.PartyLeaderChangedMessage;
import dev.emortal.api.message.party.PartyOpenChangedMessage;
import dev.emortal.api.message.party.PartyPlayerJoinedMessage;
import dev.emortal.api.message.party.PartyPlayerLeftMessage;
import dev.emortal.api.message.party.PartySettingsChangedMessage;
import dev.emortal.api.message.permission.PlayerRolesUpdateMessage;
import dev.emortal.api.message.permission.RoleUpdateMessage;
import dev.emortal.api.message.relationship.FriendAddedMessage;
import dev.emortal.api.message.relationship.FriendRemovedMessage;
import dev.emortal.api.message.relationship.FriendRequestReceivedMessage;
import org.jetbrains.annotations.NotNull;
import org.jetbrains.annotations.Nullable;

import java.util.Map;
import java.util.concurrent.ConcurrentHashMap;

public class ProtoParserRegistry {
    // Map<full descriptor name, parser>
    private static final Map<String, MessageProtoConfig<? extends Message>> parsers = new ConcurrentHashMap<>();

    static {
        registerDefaults();
    }

    public static MessageProtoConfig<? extends Message> getParser(@NotNull String fullDescriptorName) {
        return parsers.get(fullDescriptorName);
    }

    public static <T extends Message> MessageProtoConfig<T> getParser(@NotNull Class<T> clazz) {
        for (MessageProtoConfig<? extends Message> parser : parsers.values()) {
            if (parser.example().getClass().equals(clazz)) {
                //noinspection unchecked
                return (MessageProtoConfig<T>) parser;
            }
        }

        return null;
    }

    /**
     * @param fullDescriptorName The full descriptor name of the proto message.
     *                           Equal to Message.getDescriptorForType().getFullName() in Java
     *                           or Message.ProtoReflect().Descriptor().FullName() in Golang.
     * @param bytes              The bytes to parse
     * @param <T>                The type of the proto message
     * @return The parsed proto message
     * @throws ParserNotFoundException        If no parser is found for the proto message
     * @throws InvalidProtocolBufferException If the passed message (bytes) are invalid
     */
    public static <T extends Message> T parse(@NotNull String fullDescriptorName, byte[] bytes) throws ParserNotFoundException, InvalidProtocolBufferException {
        MessageProtoConfig<T> parser = (MessageProtoConfig<T>) parsers.get(fullDescriptorName);
        if (parser == null) throw new ParserNotFoundException(fullDescriptorName);

        return parser.parser().parse(bytes);
    }

    public static <T extends Message> void registerRMQ(@NotNull T example, @NotNull ProtoParser<T> parser, @Nullable String exchangeName, @Nullable String routingKey) {
        parsers.put(example.getDescriptorForType().getFullName(), new MessageProtoConfig<>(MessagingService.RABBIT_MQ, parser, example, exchangeName, routingKey, null));
    }

    public static <T extends Message> void registerRMQ(@NotNull T example, @NotNull ProtoParser<T> parser) {
        registerRMQ(example, parser, null, null);
    }

    public static <T extends Message> void registerKafka(@NotNull T example, @NotNull ProtoParser<T> parser, @NotNull String topic) {
        parsers.put(example.getDescriptorForType().getFullName(), new MessageProtoConfig<>(MessagingService.KAFKA, parser, example, null, null, topic));
    }

    private static void registerDefaults() {
        // Account Connection Manager
        registerKafka(AccountConnectedMessage.getDefaultInstance(), AccountConnectedMessage::parseFrom, "account-connections");
        registerKafka(AccountConnectionRemovedMessage.getDefaultInstance(), AccountConnectionRemovedMessage::parseFrom, "account-connections");

        // Badges
        registerKafka(PlayerBadgeAddedMessage.getDefaultInstance(), PlayerBadgeAddedMessage::parseFrom, "badge-manager");
        registerKafka(PlayerBadgeRemovedMessage.getDefaultInstance(), PlayerBadgeRemovedMessage::parseFrom, "badge-manager");
        registerKafka(PlayerActiveBadgeChangedMessage.getDefaultInstance(), PlayerActiveBadgeChangedMessage::parseFrom, "badge-manager");

        // Friend
        registerRMQ(FriendRequestReceivedMessage.getDefaultInstance(), FriendRequestReceivedMessage::parseFrom);
        registerRMQ(FriendAddedMessage.getDefaultInstance(), FriendAddedMessage::parseFrom);
        registerRMQ(FriendRemovedMessage.getDefaultInstance(), FriendRemovedMessage::parseFrom);

        // Party
        registerKafka(PartyCreatedMessage.getDefaultInstance(), PartyCreatedMessage::parseFrom, "party-manager");
        registerKafka(PartyDeletedMessage.getDefaultInstance(), PartyDeletedMessage::parseFrom, "party-manager");
        registerKafka(PartyEmptiedMessage.getDefaultInstance(), PartyEmptiedMessage::parseFrom, "party-manager");
        registerKafka(PartyOpenChangedMessage.getDefaultInstance(), PartyOpenChangedMessage::parseFrom, "party-manager");
        registerKafka(PartyInviteCreatedMessage.getDefaultInstance(), PartyInviteCreatedMessage::parseFrom, "party-manager");
        registerKafka(PartyPlayerJoinedMessage.getDefaultInstance(), PartyPlayerJoinedMessage::parseFrom, "party-manager");
        registerKafka(PartyPlayerLeftMessage.getDefaultInstance(), PartyPlayerLeftMessage::parseFrom, "party-manager");
        registerKafka(PartyLeaderChangedMessage.getDefaultInstance(), PartyLeaderChangedMessage::parseFrom, "party-manager");
        registerKafka(PartySettingsChangedMessage.getDefaultInstance(), PartySettingsChangedMessage::parseFrom, "party-manager");

        // Permission
        registerRMQ(RoleUpdateMessage.getDefaultInstance(), RoleUpdateMessage::parseFrom, "permission-manager", "role_update");
        registerRMQ(PlayerRolesUpdateMessage.getDefaultInstance(), PlayerRolesUpdateMessage::parseFrom, "permission-manager", "player_role_update");

        // Player tracker
        registerRMQ(PlayerConnectMessage.getDefaultInstance(), PlayerConnectMessage::parseFrom);
        registerRMQ(PlayerDisconnectMessage.getDefaultInstance(), PlayerDisconnectMessage::parseFrom);

        // Message handler
        registerKafka(PrivateMessageCreatedMessage.getDefaultInstance(), PrivateMessageCreatedMessage::parseFrom, "message-handler");
        registerKafka(ChatMessageCreatedMessage.getDefaultInstance(), ChatMessageCreatedMessage::parseFrom, "message-handler");

        // Common
        registerKafka(PlayerConnectMessage.getDefaultInstance(), PlayerConnectMessage::parseFrom, "mc-connections");
        registerKafka(PlayerDisconnectMessage.getDefaultInstance(), PlayerDisconnectMessage::parseFrom, "mc-connections");
        registerKafka(PlayerSwitchServerMessage.getDefaultInstance(), PlayerSwitchServerMessage::parseFrom, "mc-connections");
    }
}
