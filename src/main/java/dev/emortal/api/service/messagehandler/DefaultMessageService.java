package dev.emortal.api.service.messagehandler;

import dev.emortal.api.grpc.messagehandler.MessageHandlerGrpc;
import dev.emortal.api.grpc.messagehandler.MessageHandlerProto;
import dev.emortal.api.model.messagehandler.PrivateMessage;
import dev.emortal.api.utils.internal.GrpcErrorHelper;
import io.grpc.Channel;
import io.grpc.StatusRuntimeException;
import org.jetbrains.annotations.ApiStatus;
import org.jetbrains.annotations.NotNull;

/**
 * The default implementation of {@link MessageService} that uses a blocking stub to communicate with the gRPC server.
 */
@ApiStatus.Internal
public final class DefaultMessageService implements MessageService {

    private final MessageHandlerGrpc.MessageHandlerBlockingStub grpc;

    @ApiStatus.Internal
    public DefaultMessageService(@NotNull Channel channel) {
        this.grpc = MessageHandlerGrpc.newBlockingStub(channel);
    }

    @Override
    public @NotNull SendPrivateMessageResult sendPrivateMessage(@NotNull PrivateMessage message) {
        var request = MessageHandlerProto.PrivateMessageRequest.newBuilder().setMessage(message).build();

        MessageHandlerProto.PrivateMessageResponse response;
        try {
            response = this.grpc.sendPrivateMessage(request);
        } catch (StatusRuntimeException exception) {
            var error = GrpcErrorHelper.unpackError(exception, MessageHandlerProto.PrivateMessageErrorResponse.class);

            return switch (error.getReason()) {
                case PLAYER_NOT_ONLINE -> SendPrivateMessageResult.Error.PLAYER_NOT_ONLINE;
                case PRIVACY_BLOCKED -> SendPrivateMessageResult.Error.PRIVACY_BLOCKED;
                case YOU_BLOCKED -> SendPrivateMessageResult.Error.YOU_BLOCKED;
                case UNRECOGNIZED -> throw exception;
            };
        }

        return new SendPrivateMessageResult.Success(response.getMessage());
    }
}
