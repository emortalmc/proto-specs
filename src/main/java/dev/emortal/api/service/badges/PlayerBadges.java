package dev.emortal.api.service.badges;

import dev.emortal.api.model.badge.Badge;
import org.jetbrains.annotations.NotNull;
import org.jetbrains.annotations.Nullable;

import java.util.List;

public record PlayerBadges(@NotNull List<Badge> badges, @Nullable String activeBadgeId) {
}
