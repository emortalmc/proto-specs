package dev.emortal.api.service.mcplayer;

import dev.emortal.api.model.common.PageData;
import dev.emortal.api.model.mcplayer.LoginSession;
import org.jetbrains.annotations.NotNull;

import java.util.List;

public record LoginSessions(@NotNull List<LoginSession> sessions, @NotNull PageData pageData) {
}
