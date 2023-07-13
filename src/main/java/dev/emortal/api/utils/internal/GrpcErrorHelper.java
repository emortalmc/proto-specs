package dev.emortal.api.utils.internal;

import com.google.protobuf.InvalidProtocolBufferException;
import com.google.protobuf.Message;
import com.google.rpc.Status;
import io.grpc.StatusRuntimeException;
import io.grpc.protobuf.StatusProto;
import org.jetbrains.annotations.ApiStatus;
import org.jetbrains.annotations.NotNull;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

@ApiStatus.Internal
public final class GrpcErrorHelper {
    private static final Logger LOGGER = LoggerFactory.getLogger(GrpcErrorHelper.class);

    @ApiStatus.Internal
    public static <T extends Message> @NotNull T unpackError(@NotNull StatusRuntimeException exception, @NotNull Class<T> errorType) {
        Status status = StatusProto.fromStatusAndTrailers(exception.getStatus(), exception.getTrailers());
        if (status.getDetailsCount() == 0) {
            LOGGER.error("No error details provided in response from server!", exception);
            throw exception;
        }

        try {
            return status.getDetails(0).unpack(errorType);
        } catch (InvalidProtocolBufferException invalidException) {
            LOGGER.error("Error details provided in response from server were invalid!", exception);
            throw exception;
        }
    }

    private GrpcErrorHelper() {
    }
}
