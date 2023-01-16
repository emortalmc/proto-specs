package dev.emortal.api.utils.parser;

import com.google.protobuf.InvalidProtocolBufferException;

@FunctionalInterface
public interface ProtoParser<T> {

    T parse(byte[] bytes) throws InvalidProtocolBufferException;
}
