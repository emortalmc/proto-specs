package dev.emortal.api.utils.kafka;

import dev.emortal.api.message.common.PlayerConnectMessage;

import java.util.Random;
import java.util.UUID;
import java.util.concurrent.ExecutionException;

public class ManualTestFriendlyKafkaProducer {

    public static void main(String[] args) throws InterruptedException, ExecutionException {
        KafkaSettings settings = new KafkaSettings.Builder()
                .clientId("proto-specs-manual-test")
                .bootstrapServers("localhost:9092").build();

        FriendlyKafkaProducer producer = new FriendlyKafkaProducer(settings);

        var future = producer.produce("mc-connections", PlayerConnectMessage.newBuilder()
                .setPlayerId(UUID.randomUUID().toString())
                .setServerId("test-server-" + new Random().nextInt(1000)).build());

        future.get();
        System.out.println("Sent message");
    }
}
