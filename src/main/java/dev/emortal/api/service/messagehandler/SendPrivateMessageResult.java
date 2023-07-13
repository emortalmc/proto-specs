package dev.emortal.api.service.messagehandler;

import dev.emortal.api.model.messagehandler.PrivateMessage;
import org.jetbrains.annotations.NotNull;

public sealed interface SendPrivateMessageResult {

    record Success(@NotNull PrivateMessage message) implements SendPrivateMessageResult {
    }

    enum Error implements SendPrivateMessageResult {

        PLAYER_NOT_ONLINE,
        PRIVACY_BLOCKED,
        YOU_BLOCKED
    }
}
