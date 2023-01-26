package dev.emortal.api.utils.parser;

import com.google.protobuf.Empty;
import com.google.protobuf.InvalidProtocolBufferException;
import com.google.protobuf.Message;
import dev.emortal.api.message.common.SwitchPlayersServerMessage;
import dev.emortal.api.message.permission.PlayerRolesUpdateMessage;
import dev.emortal.api.message.permission.RoleUpdateMessage;
import dev.emortal.api.message.playertracker.PlayerConnectMessage;
import dev.emortal.api.message.playertracker.PlayerDisconnectMessage;
import dev.emortal.api.message.privatemessage.PrivateMessageReceivedMessage;
import dev.emortal.api.message.relationship.FriendAddedMessage;
import dev.emortal.api.message.relationship.FriendRemovedMessage;
import dev.emortal.api.message.relationship.FriendRequestReceivedMessage;
import org.jetbrains.annotations.NotNull;

import java.util.Map;
import java.util.concurrent.ConcurrentHashMap;

public class ProtoParserRegistry {
    // Map<full descriptor name, parser>
    private static final Map<String, ProtoParser<?>> parsers = new ConcurrentHashMap<>();

    static {
        registerDefaults();
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
    public static <T> T parse(@NotNull String fullDescriptorName, byte[] bytes) throws ParserNotFoundException, InvalidProtocolBufferException {
        ProtoParser<T> parser = (ProtoParser<T>) parsers.get(fullDescriptorName);
        if (parser == null) throw new ParserNotFoundException(fullDescriptorName);

        return parser.parse(bytes);
    }

    public static <T extends Message> void register(@NotNull T example, @NotNull ProtoParser<T> parser) {
        parsers.put(example.getDescriptorForType().getFullName(), parser);
    }

    private static void registerDefaults() {
        register(Empty.getDefaultInstance(), Empty::parseFrom); // Mostly as an example. Empty shouldn't need to be parsed.

        // Friend
        register(FriendRequestReceivedMessage.getDefaultInstance(), FriendRequestReceivedMessage::parseFrom);
        register(FriendAddedMessage.getDefaultInstance(), FriendAddedMessage::parseFrom);
        register(FriendRemovedMessage.getDefaultInstance(), FriendRemovedMessage::parseFrom);

        // Permission
        register(RoleUpdateMessage.getDefaultInstance(), RoleUpdateMessage::parseFrom);
        register(PlayerRolesUpdateMessage.getDefaultInstance(), PlayerRolesUpdateMessage::parseFrom);

        // Player tracker
        register(PlayerConnectMessage.getDefaultInstance(), PlayerConnectMessage::parseFrom);
        register(PlayerDisconnectMessage.getDefaultInstance(), PlayerDisconnectMessage::parseFrom);

        // Private message
        register(PrivateMessageReceivedMessage.getDefaultInstance(), PrivateMessageReceivedMessage::parseFrom);

        // Common
        register(SwitchPlayersServerMessage.getDefaultInstance(), SwitchPlayersServerMessage::parseFrom);
    }
}
