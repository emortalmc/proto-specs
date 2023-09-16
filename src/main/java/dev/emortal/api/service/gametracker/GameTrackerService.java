package dev.emortal.api.service.gametracker;

import org.jetbrains.annotations.Blocking;

/**
 * A service interface for the game tracker service.
 *
 * <p>
 * All methods within this class are blocking. It is up to the client to call these methods asynchronously.
 * <br>
 * The recommended way to do this is to use {@linkplain Thread#startVirtualThread virtual threads}.
 */
@Blocking
public interface GameTrackerService {
}
