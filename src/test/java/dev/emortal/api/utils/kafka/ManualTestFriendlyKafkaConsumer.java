package dev.emortal.api.utils.kafka;

import dev.emortal.api.message.common.PlayerConnectMessage;

import java.util.concurrent.CountDownLatch;

public final class ManualTestFriendlyKafkaConsumer {

    public static void main(String[] args) throws InterruptedException {
        var settings = KafkaSettings.builder()
                .clientId("proto-specs-manual-test")
                .bootstrapServers("localhost:9092")
                .build();

        var consumer = new FriendlyKafkaConsumer(settings);

        var latch = new CountDownLatch(1);
        consumer.addListener(PlayerConnectMessage.class, message -> {
            System.out.println("Received message: " + message);
            latch.countDown();
        });

        latch.await();
    }
}
