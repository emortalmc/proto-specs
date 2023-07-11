package dev.emortal.api.service.badges;

import dev.emortal.api.model.badge.Badge;
import org.jetbrains.annotations.NotNull;
import org.jetbrains.annotations.Nullable;

import java.util.List;
import java.util.UUID;

public interface BadgeService {

    @NotNull List<Badge> getAllBadges();

    @NotNull PlayerBadges getPlayerBadges(@NotNull UUID playerId);

    @Nullable Badge getActiveBadge(@NotNull UUID playerId);

    @NotNull SetActiveBadgeResult setActiveBadge(@NotNull UUID playerId, @NotNull String badgeId);

    @NotNull AddBadgeToPlayerResult addBadgeToPlayer(@NotNull UUID playerId, @NotNull String badgeId);

    @NotNull RemoveBadgeFromPlayerResult removeBadgeFromPlayer(@NotNull UUID playerId, @NotNull String badgeId);
}
