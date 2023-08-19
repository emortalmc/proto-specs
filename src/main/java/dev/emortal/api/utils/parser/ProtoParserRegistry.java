package dev.emortal.api.utils.parser;

import com.google.protobuf.InvalidProtocolBufferException;
import com.google.protobuf.Message;
import dev.emortal.api.message.accountconnmanager.AccountConnectedMessage;
import dev.emortal.api.message.accountconnmanager.AccountConnectionRemovedMessage;
import dev.emortal.api.message.badge.PlayerActiveBadgeChangedMessage;
import dev.emortal.api.message.badge.PlayerBadgeAddedMessage;
import dev.emortal.api.message.badge.PlayerBadgeRemovedMessage;
import dev.emortal.api.message.common.PlayerChatMessageMessage;
import dev.emortal.api.message.common.PlayerConnectMessage;
import dev.emortal.api.message.common.PlayerDisconnectMessage;
import dev.emortal.api.message.common.PlayerSwitchServerMessage;
import dev.emortal.api.message.gamedata.UpdateGamePlayerDataMessage;
import dev.emortal.api.message.gamesdk.GameReadyMessage;
import dev.emortal.api.message.matchmaker.MatchCreatedMessage;
import dev.emortal.api.message.matchmaker.PendingMatchCreatedMessage;
import dev.emortal.api.message.matchmaker.PendingMatchDeletedMessage;
import dev.emortal.api.message.matchmaker.PendingMatchUpdatedMessage;
import dev.emortal.api.message.matchmaker.TicketCreatedMessage;
import dev.emortal.api.message.matchmaker.TicketDeletedMessage;
import dev.emortal.api.message.matchmaker.TicketUpdatedMessage;
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
import dev.emortal.api.message.relationship.FriendConnectionMessage;
import dev.emortal.api.message.relationship.FriendRemovedMessage;
import dev.emortal.api.message.relationship.FriendRequestReceivedMessage;
import org.jetbrains.annotations.NotNull;
import org.jetbrains.annotations.Nullable;

import java.util.Map;
import java.util.concurrent.ConcurrentHashMap;

public final class ProtoParserRegistry {
    // Map<full descriptor name, parser>
    private static final Map<String, MessageProtoConfig<? extends Message>> parsers = new ConcurrentHashMap<>();

    static {
        registerDefaults();
    }

    public static @Nullable MessageProtoConfig<? extends Message> getParser(@NotNull String fullDescriptorName) {
        return parsers.get(fullDescriptorName);
    }

    public static <T extends Message> @Nullable MessageProtoConfig<T> getParser(@NotNull Class<T> clazz) {
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
    public static <T extends Message> @NotNull T parse(@NotNull String fullDescriptorName,
                                                       byte[] bytes) throws ParserNotFoundException, InvalidProtocolBufferException {
        @SuppressWarnings("unchecked") // This is safe here because the register method enforces this guarantee
        MessageProtoConfig<T> parser = (MessageProtoConfig<T>) parsers.get(fullDescriptorName);
        if (parser == null) throw new ParserNotFoundException(fullDescriptorName);

        return parser.parser().parse(bytes);
    }

    public static <T extends Message> void register(@NotNull T example, @NotNull ProtoParser<T> parser, @NotNull String topic) {
        parsers.put(example.getDescriptorForType().getFullName(), new MessageProtoConfig<>(parser, example, topic));
    }

    private static void registerDefaults() {
        // Account Connection Manager
        register(AccountConnectedMessage.getDefaultInstance(), AccountConnectedMessage::parseFrom, "account-connections");
        register(AccountConnectionRemovedMessage.getDefaultInstance(), AccountConnectionRemovedMessage::parseFrom, "account-connections");

        // Badges
        register(PlayerBadgeAddedMessage.getDefaultInstance(), PlayerBadgeAddedMessage::parseFrom, "badge-manager");
        register(PlayerBadgeRemovedMessage.getDefaultInstance(), PlayerBadgeRemovedMessage::parseFrom, "badge-manager");
        register(PlayerActiveBadgeChangedMessage.getDefaultInstance(), PlayerActiveBadgeChangedMessage::parseFrom, "badge-manager");

        // Game Player Data
        register(UpdateGamePlayerDataMessage.getDefaultInstance(), UpdateGamePlayerDataMessage::parseFrom, "game-player-data");

        // (TODO) Game Tracker

        // Friend
        register(FriendRequestReceivedMessage.getDefaultInstance(), FriendRequestReceivedMessage::parseFrom, "relationship-manager");
        register(FriendAddedMessage.getDefaultInstance(), FriendAddedMessage::parseFrom, "relationship-manager");
        register(FriendRemovedMessage.getDefaultInstance(), FriendRemovedMessage::parseFrom, "relationship-manager");
        register(FriendConnectionMessage.getDefaultInstance(), FriendConnectionMessage::parseFrom, "relationship-manager");

        // Party
        register(PartyCreatedMessage.getDefaultInstance(), PartyCreatedMessage::parseFrom, "party-manager");
        register(PartyDeletedMessage.getDefaultInstance(), PartyDeletedMessage::parseFrom, "party-manager");
        register(PartyEmptiedMessage.getDefaultInstance(), PartyEmptiedMessage::parseFrom, "party-manager");
        register(PartyOpenChangedMessage.getDefaultInstance(), PartyOpenChangedMessage::parseFrom, "party-manager");
        register(PartyInviteCreatedMessage.getDefaultInstance(), PartyInviteCreatedMessage::parseFrom, "party-manager");
        register(PartyPlayerJoinedMessage.getDefaultInstance(), PartyPlayerJoinedMessage::parseFrom, "party-manager");
        register(PartyPlayerLeftMessage.getDefaultInstance(), PartyPlayerLeftMessage::parseFrom, "party-manager");
        register(PartyLeaderChangedMessage.getDefaultInstance(), PartyLeaderChangedMessage::parseFrom, "party-manager");
        register(PartySettingsChangedMessage.getDefaultInstance(), PartySettingsChangedMessage::parseFrom, "party-manager");

        // Permission
        register(RoleUpdateMessage.getDefaultInstance(), RoleUpdateMessage::parseFrom, "permission-manager");
        register(PlayerRolesUpdateMessage.getDefaultInstance(), PlayerRolesUpdateMessage::parseFrom, "permission-manager");

        // Message handler
        register(PrivateMessageCreatedMessage.getDefaultInstance(), PrivateMessageCreatedMessage::parseFrom, "message-handler");
        register(ChatMessageCreatedMessage.getDefaultInstance(), ChatMessageCreatedMessage::parseFrom, "message-handler");

        // Matchmaker
        register(TicketCreatedMessage.getDefaultInstance(), TicketCreatedMessage::parseFrom, "matchmaker");
        register(TicketUpdatedMessage.getDefaultInstance(), TicketUpdatedMessage::parseFrom, "matchmaker");
        register(TicketDeletedMessage.getDefaultInstance(), TicketDeletedMessage::parseFrom, "matchmaker");
        register(PendingMatchCreatedMessage.getDefaultInstance(), PendingMatchCreatedMessage::parseFrom, "matchmaker");
        register(PendingMatchUpdatedMessage.getDefaultInstance(), PendingMatchUpdatedMessage::parseFrom, "matchmaker");
        register(PendingMatchDeletedMessage.getDefaultInstance(), PendingMatchDeletedMessage::parseFrom, "matchmaker");
        register(MatchCreatedMessage.getDefaultInstance(), MatchCreatedMessage::parseFrom, "matchmaker");

        // Game SDK
        register(GameReadyMessage.getDefaultInstance(), GameReadyMessage::parseFrom, "game-sdk");

        // Common
        register(PlayerConnectMessage.getDefaultInstance(), PlayerConnectMessage::parseFrom, "mc-connections");
        register(PlayerDisconnectMessage.getDefaultInstance(), PlayerDisconnectMessage::parseFrom, "mc-connections");
        register(PlayerSwitchServerMessage.getDefaultInstance(), PlayerSwitchServerMessage::parseFrom, "mc-connections");
        register(PlayerChatMessageMessage.getDefaultInstance(), PlayerChatMessageMessage::parseFrom, "mc-messages");
    }

    private ProtoParserRegistry() {
    }
}
