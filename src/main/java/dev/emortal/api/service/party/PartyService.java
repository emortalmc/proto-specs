package dev.emortal.api.service.party;

import dev.emortal.api.model.common.Pageable;
import dev.emortal.api.model.common.PlayerSkin;
import dev.emortal.api.model.party.EventData;
import dev.emortal.api.model.party.Party;
import dev.emortal.api.model.party.PartyInvite;
import org.jetbrains.annotations.Blocking;
import org.jetbrains.annotations.NotNull;
import org.jetbrains.annotations.Nullable;

import java.time.Instant;
import java.util.List;
import java.util.UUID;

/**
 * A service interface for the party service.
 *
 * <p>
 * All methods within this class are blocking. It is up to the client to call these methods asynchronously.
 * <br>
 * The recommended way to do this is to use {@linkplain Thread#startVirtualThread virtual threads}.
 */
@Blocking
public interface PartyService {

    @Nullable Party getParty(@NotNull String partyId);

    @Nullable Party getPartyByPlayer(@NotNull UUID playerId);

    @NotNull ModifyPartyResult emptyParty(@NotNull String partyId);

    @NotNull ModifyPartyResult emptyPartyByPlayer(@NotNull UUID playerId);

    @NotNull ModifyPartyResult setPartyOpen(@NotNull UUID playerId, boolean open);

    @NotNull List<PartyInvite> getInvites(@NotNull String partyId, @NotNull Pageable pageable);

    @NotNull List<PartyInvite> getInvitesByPlayer(@NotNull UUID playerId, @NotNull Pageable pageable);

    @NotNull InvitePlayerToPartyResult invitePlayer(@NotNull UUID inviterId, @NotNull String inviterName,
                                                    @NotNull UUID invitedId, @NotNull String invitedName);

    @NotNull JoinPartyResult joinParty(@NotNull UUID joinerId, @NotNull String joinerName, @NotNull UUID targetId);

    @NotNull LeavePartyResult leaveParty(@NotNull UUID leaverId);

    @NotNull KickPlayerFromPartyResult kickPlayer(@NotNull UUID kickerId, @NotNull String kickerName, @NotNull UUID targetId);

    @NotNull SetPartyLeaderResult setPartyLeader(@NotNull UUID requesterId, @NotNull String requesterName, @NotNull UUID newLeaderId);

    @NotNull EventData createEvent(@NotNull String id, @NotNull UUID ownerId, @NotNull String ownerUsername, @NotNull PlayerSkin skin,
                                   @Nullable Instant showTime, @Nullable Instant startTime);

    @NotNull EventData updateEvent(@NotNull String id, @Nullable Instant showTime, @Nullable Instant startTime);

    @NotNull List<EventData> listEvents();

    /**
     * Deletes an event by its id or if the id is null, deletes the current ongoing event.
     *
     * @param id the id of the event.
     */
    DeleteEventResult deleteEvent(@Nullable String id);
}
