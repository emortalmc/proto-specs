package dev.emortal.api.utils.parser;

import com.google.protobuf.Empty;
import com.google.protobuf.InvalidProtocolBufferException;
import com.google.protobuf.Message;
import dev.emortal.api.message.common.PlayerConnectMessage;
import dev.emortal.api.message.common.PlayerDisconnectMessage;
import dev.emortal.api.message.common.SwitchPlayersServerMessage;
import dev.emortal.api.message.party.PartyCreatedMessage;
import dev.emortal.api.message.party.PartyDisbandedMessage;
import dev.emortal.api.message.party.PartyInviteCreatedMessage;
import dev.emortal.api.message.party.PartyLeaderChangedMessage;
import dev.emortal.api.message.party.PartyPlayerJoinedMessage;
import dev.emortal.api.message.party.PartyPlayerLeftMessage;
import dev.emortal.api.message.party.PartySettingsChangedMessage;
import dev.emortal.api.message.permission.PlayerRolesUpdateMessage;
import dev.emortal.api.message.permission.RoleUpdateMessage;
import dev.emortal.api.message.privatemessage.PrivateMessageReceivedMessage;
import dev.emortal.api.message.relationship.FriendAddedMessage;
import dev.emortal.api.message.relationship.FriendRemovedMessage;
import dev.emortal.api.message.relationship.FriendRequestReceivedMessage;
import org.jetbrains.annotations.NotNull;
import org.jetbrains.annotations.Nullable;

import java.util.Map;
import java.util.concurrent.ConcurrentHashMap;

public class ProtoParserRegistry {
    // Map<full descriptor name, parser>
    private static final Map<String, ParsableProto<? extends Message>> parsers = new ConcurrentHashMap<>();

    static {
        registerDefaults();
    }

    public static ParsableProto<? extends Message> getParser(@NotNull String fullDescriptorName) {
        return parsers.get(fullDescriptorName);
    }

    public static <T extends Message> ParsableProto<T> getParser(@NotNull Class<T> clazz) {
        for (ParsableProto<? extends Message> parser : parsers.values()) {
            if (parser.example().getClass().equals(clazz)) {
                //noinspection unchecked
                return (ParsableProto<T>) parser;
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
        ParsableProto<T> parser = (ParsableProto<T>) parsers.get(fullDescriptorName);
        if (parser == null) throw new ParserNotFoundException(fullDescriptorName);

        return parser.parser().parse(bytes);
    }

    public static <T extends Message> void register(@NotNull T example, @NotNull ProtoParser<T> parser, @Nullable String exchangeName, @Nullable String routingKey) {
        parsers.put(example.getDescriptorForType().getFullName(), new ParsableProto<>(parser, example, exchangeName, routingKey));
    }

    public static <T extends Message> void register(@NotNull T example, @NotNull ProtoParser<T> parser) {
        register(example, parser, null, null);
    }

    private static void registerDefaults() {
        register(Empty.getDefaultInstance(), Empty::parseFrom); // Mostly as an example. Empty shouldn't need to be parsed.

        // Friend
        register(FriendRequestReceivedMessage.getDefaultInstance(), FriendRequestReceivedMessage::parseFrom);
        register(FriendAddedMessage.getDefaultInstance(), FriendAddedMessage::parseFrom);
        register(FriendRemovedMessage.getDefaultInstance(), FriendRemovedMessage::parseFrom);

        // Party
        register(PartyCreatedMessage.getDefaultInstance(), PartyCreatedMessage::parseFrom, "party-manager", "party_created");
        register(PartyDisbandedMessage.getDefaultInstance(), PartyDisbandedMessage::parseFrom, "party-manager", "party_disbanded");
        register(PartyInviteCreatedMessage.getDefaultInstance(), PartyInviteCreatedMessage::parseFrom, "party-manager", "party_invite_created");
        register(PartyPlayerJoinedMessage.getDefaultInstance(), PartyPlayerJoinedMessage::parseFrom, "party-manager", "party_player_joined");
        register(PartyPlayerLeftMessage.getDefaultInstance(), PartyPlayerLeftMessage::parseFrom, "party-manager", "party_player_left");
        register(PartyLeaderChangedMessage.getDefaultInstance(), PartyLeaderChangedMessage::parseFrom, "party-manager", "party_leader_changed");
        register(PartySettingsChangedMessage.getDefaultInstance(), PartySettingsChangedMessage::parseFrom, "party-manager", "party_settings_changed");

        // Permission
        register(RoleUpdateMessage.getDefaultInstance(), RoleUpdateMessage::parseFrom, "permission-manager", "role_update");
        register(PlayerRolesUpdateMessage.getDefaultInstance(), PlayerRolesUpdateMessage::parseFrom, "permission-manager", "player_role_update");

        // Player tracker
        register(PlayerConnectMessage.getDefaultInstance(), PlayerConnectMessage::parseFrom);
        register(PlayerDisconnectMessage.getDefaultInstance(), PlayerDisconnectMessage::parseFrom);

        // Private message
        register(PrivateMessageReceivedMessage.getDefaultInstance(), PrivateMessageReceivedMessage::parseFrom);

        // Common
        register(SwitchPlayersServerMessage.getDefaultInstance(), SwitchPlayersServerMessage::parseFrom);
    }
}
