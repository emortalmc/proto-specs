package dev.emortal.api.service.party;

import dev.emortal.api.model.party.PartyInvite;
import org.jetbrains.annotations.NotNull;

public sealed interface InvitePlayerToPartyResult {

    record Success(@NotNull PartyInvite invite) implements InvitePlayerToPartyResult {
    }

    enum Error implements InvitePlayerToPartyResult {

        TARGET_IN_THIS_PARTY,
        TARGET_IN_OTHER_PARTY,
        TARGET_ALREADY_INVITED,
        NO_PERMISSION,
        PARTY_OPEN
    }
}
