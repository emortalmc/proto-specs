package dev.emortal.api.service.badges;

import dev.emortal.api.model.badge.Badge;
import org.jetbrains.annotations.Blocking;
import org.jetbrains.annotations.NotNull;
import org.jetbrains.annotations.Nullable;

import java.util.List;
import java.util.UUID;

/**
 * A service interface for the badge service.
 *
 * <p>
 * All methods within this class are blocking. It is up to the client to call these methods asynchronously.
 * <br>
 * The recommended way to do this is to use {@linkplain Thread#startVirtualThread virtual threads}.
 */
@Blocking
public interface BadgeService {

    @NotNull List<Badge> getAllBadges();

    @NotNull PlayerBadges getPlayerBadges(@NotNull UUID playerId);

    @Nullable Badge getActiveBadge(@NotNull UUID playerId);

    @NotNull SetActiveBadgeResult setActiveBadge(@NotNull UUID playerId, @NotNull String badgeId);

    @NotNull AddBadgeToPlayerResult addBadgeToPlayer(@NotNull UUID playerId, @NotNull String badgeId);

    @NotNull RemoveBadgeFromPlayerResult removeBadgeFromPlayer(@NotNull UUID playerId, @NotNull String badgeId);
}
