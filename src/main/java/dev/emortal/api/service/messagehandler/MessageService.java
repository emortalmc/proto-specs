package dev.emortal.api.service.messagehandler;

import dev.emortal.api.model.messagehandler.PrivateMessage;
import org.jetbrains.annotations.Blocking;
import org.jetbrains.annotations.NotNull;

/**
 * A service interface for the message handler.
 *
 * <p>
 * All methods within this class are blocking. It is up to the client to call these methods asynchronously.
 * <br>
 * The recommended way to do this is to use {@linkplain Thread#startVirtualThread virtual threads}.
 */
@Blocking
public interface MessageService {

    @NotNull SendPrivateMessageResult sendPrivateMessage(@NotNull PrivateMessage message);
}
