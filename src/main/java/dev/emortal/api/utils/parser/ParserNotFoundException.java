package dev.emortal.api.utils.parser;

import org.jetbrains.annotations.NotNull;

@SuppressWarnings("unused")
public class ParserNotFoundException extends RuntimeException {
    private final @NotNull String fullDescriptorName;

    ParserNotFoundException(@NotNull String fullDescriptorName) {
        super("parser not found for proto " + fullDescriptorName);
        this.fullDescriptorName = fullDescriptorName;
    }

    public @NotNull String getFullDescriptorName() {
        return fullDescriptorName;
    }
}
