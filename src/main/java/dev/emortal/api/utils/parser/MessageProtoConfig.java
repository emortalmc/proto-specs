package dev.emortal.api.utils.parser;

import com.google.protobuf.Message;
import org.jetbrains.annotations.NotNull;

/**
 * Contains a parser for the message type.
 */
public record MessageProtoConfig<T extends Message>(@NotNull ProtoParser<T> parser, @NotNull T example, @NotNull String topic) {

}
