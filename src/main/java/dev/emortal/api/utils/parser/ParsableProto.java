package dev.emortal.api.utils.parser;

import com.google.protobuf.Message;
import org.jetbrains.annotations.NotNull;
import org.jetbrains.annotations.Nullable;

/**
 * Contains a parser for the message type.
 */
public record ParsableProto<T extends Message>(@NotNull ProtoParser<T> parser, @Nullable String exchangeName,
                                               @Nullable String routingKey) {

}
