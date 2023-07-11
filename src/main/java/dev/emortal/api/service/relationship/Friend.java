package dev.emortal.api.service.relationship;

import org.jetbrains.annotations.NotNull;

import java.time.Instant;
import java.util.UUID;

public record Friend(@NotNull UUID id, @NotNull Instant friendsSince) {
}
