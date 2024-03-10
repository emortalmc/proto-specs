package dev.emortal.api.service.party;

public enum DeleteEventResult {
    SUCCESS,
    NOT_FOUND,
    /**
     * Sent if the ID was not provided and there is no current event
     */
    NO_CURRENT_EVENT
}
