package dev.emortal.api.service.relationship;

import org.jetbrains.annotations.NotNull;

import java.time.Instant;
import java.util.UUID;

public record RequestedFriend(@NotNull UUID requesterId, @NotNull UUID targetId, @NotNull Instant requestTime) {
}
