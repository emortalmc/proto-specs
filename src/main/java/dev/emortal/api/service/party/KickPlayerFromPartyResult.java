package dev.emortal.api.service.party;

public enum KickPlayerFromPartyResult {

    SUCCESS,
    TARGET_NOT_IN_PARTY,
    SELF_NOT_LEADER,
    TARGET_IS_LEADER
}
