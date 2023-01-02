package dev.emortal.api.utils.spring.autoconfiguration;

import dev.emortal.api.utils.spring.exception.DefaultGrpcExceptionHandler;
import org.springframework.boot.autoconfigure.AutoConfiguration;
import org.springframework.context.annotation.Bean;

@AutoConfiguration
public class GrpcExceptionHandlerAutoConfiguration {

    @Bean
    public DefaultGrpcExceptionHandler defaultGrpcExceptionHandler() {
        return new DefaultGrpcExceptionHandler();
    }
}
