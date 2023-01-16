package dev.emortal.api.utils.parser;

public class ParserNotFoundException extends RuntimeException {

    ParserNotFoundException(String fullDescriptorName) {
        super("parser not found for proto " + fullDescriptorName);
    }

}
