package dev.emortal.api.service.messagehandler;

import dev.emortal.api.model.messagehandler.PrivateMessage;
import org.jetbrains.annotations.NotNull;

public interface MessageService {

    @NotNull SendPrivateMessageResult sendPrivateMessage(@NotNull PrivateMessage message);
}
