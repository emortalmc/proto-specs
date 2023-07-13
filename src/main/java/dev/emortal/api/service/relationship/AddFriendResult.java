package dev.emortal.api.service.relationship;

import org.jetbrains.annotations.NotNull;

import java.time.Instant;

public sealed interface AddFriendResult {

    record RequestSent() implements AddFriendResult {
    }

    record FriendAdded(@NotNull Instant friendsSince) implements AddFriendResult {
    }

    enum Error implements AddFriendResult {

        ALREADY_FRIENDS,
        ALREADY_REQUESTED,
        YOU_BLOCKED,
        PRIVACY_BLOCKED
    }
}
