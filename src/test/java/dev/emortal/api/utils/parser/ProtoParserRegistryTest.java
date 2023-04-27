package dev.emortal.api.utils.parser;

import com.google.protobuf.InvalidProtocolBufferException;
import com.google.protobuf.Message;
import com.google.protobuf.Timestamp;
import dev.emortal.api.message.permission.RoleUpdateMessage;
import dev.emortal.api.message.relationship.FriendAddedMessage;
import dev.emortal.api.message.relationship.FriendRequestReceivedMessage;
import dev.emortal.api.model.permission.PermissionNode;
import dev.emortal.api.model.permission.Role;
import dev.emortal.api.model.relationship.FriendRequest;
import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.params.ParameterizedTest;
import org.junit.jupiter.params.provider.Arguments;
import org.junit.jupiter.params.provider.MethodSource;

import java.time.Instant;
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
        ProtoParserRegistry.registerKafka(Timestamp.getDefaultInstance(), Timestamp::parseFrom, "test");

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
        var one = FriendRequestReceivedMessage.newBuilder()
                .setRequest(
                        FriendRequest.newBuilder()
                                .setTargetId("testRecipientId")
                                .setSenderId("testSenderId")
                                .setSenderUsername("testUsername")
                )
                .build();

        var two = FriendAddedMessage.newBuilder()
                .setRecipientId("testRecipientId")
                .setSenderId("testSenderId")
                .setSenderUsername("testUsername")
                .build();

        // This message contains some fields not set on purpose
        var three = RoleUpdateMessage.newBuilder()
                .setChangeType(RoleUpdateMessage.ChangeType.MODIFY)
                .setRole(Role.newBuilder()
                        .setId("testRoleId")
                        .setDisplayName("testDisplayName")
                        .addPermissions(
                                PermissionNode.newBuilder()
                                        .setState(PermissionNode.PermissionState.ALLOW)
                                        .setNode("testNode")
                        )
                ).build();

        return Stream.of(
                Arguments.of(one.toByteArray(), one),
                Arguments.of(two.toByteArray(), two),
                Arguments.of(three.toByteArray(), three)
        );
    }
}
