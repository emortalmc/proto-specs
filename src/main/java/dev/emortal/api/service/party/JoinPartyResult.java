package dev.emortal.api.service.party;

import dev.emortal.api.model.party.Party;
import org.jetbrains.annotations.NotNull;

public sealed interface JoinPartyResult {

    record Success(@NotNull Party party) implements JoinPartyResult {
    }

    enum Error implements JoinPartyResult {

        ALREADY_IN_PARTY,
        NOT_INVITED
    }
}
