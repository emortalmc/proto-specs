package dev.emortal.api.utils.parser;

import com.google.protobuf.Empty;
import com.google.protobuf.InvalidProtocolBufferException;
import com.google.protobuf.Message;
import dev.emortal.api.messaging.friend.ProxyFriendAddedMessage;
import dev.emortal.api.messaging.friend.ProxyFriendRemovedMessage;
import dev.emortal.api.messaging.friend.ProxyFriendRequestReceivedMessage;
import dev.emortal.api.messaging.general.ProxyServerSwitchMessage;
import dev.emortal.api.messaging.privatemessage.ProxyPrivateMessageReceivedMessage;
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
     * fullDescriptorName is equal to Message.ProtoReflect().Descriptor().FullName() in Golang
     * and Message.getDescriptorForType().getFullName() in Java.
     *
     * @param fullDescriptorName The full descriptor name of the proto message
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

        register(ProxyFriendRequestReceivedMessage.getDefaultInstance(), ProxyFriendRequestReceivedMessage::parseFrom);
        register(ProxyFriendAddedMessage.getDefaultInstance(), ProxyFriendAddedMessage::parseFrom);
        register(ProxyFriendRemovedMessage.getDefaultInstance(), ProxyFriendRemovedMessage::parseFrom);

        register(ProxyServerSwitchMessage.getDefaultInstance(), ProxyServerSwitchMessage::parseFrom);

        register(ProxyPrivateMessageReceivedMessage.getDefaultInstance(), ProxyPrivateMessageReceivedMessage::parseFrom);
    }
}
