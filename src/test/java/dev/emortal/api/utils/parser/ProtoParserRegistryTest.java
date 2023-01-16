package dev.emortal.api.utils.parser;

import com.google.protobuf.InvalidProtocolBufferException;
import com.google.protobuf.Message;
import com.google.protobuf.Timestamp;
import dev.emortal.api.messaging.friend.ProxyFriendAddedMessage;
import dev.emortal.api.messaging.friend.ProxyFriendRequestReceivedMessage;
import dev.emortal.api.messaging.general.ProxyServerSwitchMessage;
import dev.emortal.api.service.ServerDiscoveryProto;
import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.params.ParameterizedTest;
import org.junit.jupiter.params.provider.Arguments;
import org.junit.jupiter.params.provider.MethodSource;

import java.time.Instant;
import java.util.Random;
import java.util.Set;
import java.util.UUID;
import java.util.stream.Stream;

public class ProtoParserRegistryTest {

    @ParameterizedTest
    @MethodSource("testParsingParams")
    public void testParsing(byte[] bytes, Message expected) throws InvalidProtocolBufferException {
        Message actual = ProtoParserRegistry.parse(expected.getDescriptorForType().getFullName(), bytes);
        Assertions.assertEquals(expected, actual);
    }

    @Test
    public void testRegisteredParsing() throws InvalidProtocolBufferException {
        ProtoParserRegistry.register(Timestamp.getDefaultInstance(), Timestamp::parseFrom);

        Instant instant = Instant.now();

        Timestamp timestamp = Timestamp.newBuilder()
                .setSeconds(instant.getEpochSecond())
                .setNanos(instant.getNano())
                .build();

        Timestamp parsedTimestamp = ProtoParserRegistry.parse(Timestamp.getDescriptor().getFullName(), timestamp.toByteArray());
        Instant parsedInstant = Instant.ofEpochSecond(parsedTimestamp.getSeconds(), parsedTimestamp.getNanos());

        Assertions.assertEquals(instant, parsedInstant);
    }

    private static Stream<Arguments> testParsingParams() {
        ProxyFriendRequestReceivedMessage one = ProxyFriendRequestReceivedMessage.newBuilder()
                .setRecipientId("test-recipient-id")
                .setSenderId("test-sender-id")
                .setSenderUsername("test-username")
                .build();

        ProxyFriendAddedMessage two = ProxyFriendAddedMessage.newBuilder()
                .setRecipientId("testRecipientId")
                .setSenderId("testSenderId")
                .setSenderUsername("testUsername")
                .build();

        ProxyServerSwitchMessage three = ProxyServerSwitchMessage.newBuilder()
                .setServer(
                        ServerDiscoveryProto.ConnectableServer.newBuilder()
                                .setId("lobby-sdgsd-32zr56")
                                .setAddress("127.0.0.1")
                                .setPort(new Random().nextInt(1, 65_535 + 1)).build()
                )
                .addAllPlayerIds(Set.of(UUID.randomUUID().toString(), UUID.randomUUID().toString(), UUID.randomUUID().toString()))
                .build();

        return Stream.of(
                Arguments.of(one.toByteArray(), one),
                Arguments.of(two.toByteArray(), two),
                Arguments.of(three.toByteArray(), three)
        );
    }
}
