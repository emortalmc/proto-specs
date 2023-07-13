package dev.emortal.api.utils.parser;

import com.google.protobuf.InvalidProtocolBufferException;
import org.jetbrains.annotations.NotNull;

@FunctionalInterface
public interface ProtoParser<T> {

    @NotNull T parse(byte[] bytes) throws InvalidProtocolBufferException;
}
