package dev.emortal.api.utils.kafka;

import dev.emortal.api.message.common.PlayerConnectMessage;

import java.util.concurrent.CountDownLatch;

public class ManualTestFriendlyKafkaConsumer {

    public static void main(String[] args) throws InterruptedException {
        KafkaSettings settings = KafkaSettings.builder()
                .clientId("proto-specs-manual-test")
                .bootstrapServers("localhost:9092").build();

        FriendlyKafkaConsumer consumer = new FriendlyKafkaConsumer(settings);

        CountDownLatch latch = new CountDownLatch(1);
        consumer.addListener(PlayerConnectMessage.class, message -> {
            System.out.println("Received message: " + message);
            latch.countDown();
        });

        latch.await();
    }
}
