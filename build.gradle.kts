import com.google.protobuf.gradle.*

plugins {
    `java-library`
    `maven-publish`

    // gRPC
    id("com.google.protobuf") version "0.9.4"
}

group = "dev.emortal.api"
version = "1.0"

val grpcVersion = "1.66.0"
val protobufVersion = "4.28.0"
val protocVersion = protobufVersion

repositories {
    mavenCentral()

    maven("https://packages.confluent.io/maven/")
}

dependencies {
    api("com.google.protobuf:protobuf-java:$protobufVersion")
    api("com.google.protobuf:protobuf-java-util:$protobufVersion")
    api("org.apache.tomcat:tomcat-annotations-api:10.1.23")
    api("org.slf4j:slf4j-api:2.0.17")

    implementation("io.micrometer:micrometer-core:1.12.5")
    implementation("org.apache.kafka:kafka-clients:7.6.0-ccs")

    compileOnly("org.jetbrains:annotations:26.0.2-1")
    annotationProcessor("org.jetbrains:annotations:26.0.2-1")

    // Protobuf uses @javax.annotation.Generated on generated types, and sometimes this isn't available.
    // This make sure it's always available
    compileOnly("javax.annotation:javax.annotation-api:1.3.2")

    implementation("com.github.ben-manes.caffeine:caffeine:3.2.2")

    api("io.grpc:grpc-netty:$grpcVersion")
    api("io.grpc:grpc-stub:$grpcVersion")
    api("io.grpc:grpc-protobuf:$grpcVersion")

    testImplementation("org.junit.jupiter:junit-jupiter-api:6.0.0-RC3")
    testImplementation("org.junit.jupiter:junit-jupiter-params:6.0.0-RC3")
    testRuntimeOnly("org.junit.platform:junit-platform-launcher:6.0.0-RC3")
    testRuntimeOnly("org.junit.jupiter:junit-jupiter-engine:6.0.0-RC3")
    testRuntimeOnly("org.slf4j:slf4j-simple:2.0.17")
}

protobuf {
    protoc {
        artifact = "com.google.protobuf:protoc:$protocVersion"
    }
    plugins {
        id("grpc") {
            artifact = "io.grpc:protoc-gen-grpc-java:$grpcVersion"
        }
    }
    generateProtoTasks {
        all().forEach {
            it.plugins {
                id("grpc") {}
            }
        }
    }
}

java {
    toolchain {
        languageVersion.set(JavaLanguageVersion.of(25))
    }

    withSourcesJar()
}

tasks {
    withType<Jar> { duplicatesStrategy = DuplicatesStrategy.EXCLUDE }
    test {
        useJUnitPlatform()
    }
}

publishing {
    repositories {
        maven {
            name = "development"
            url = uri("https://repo.emortal.dev/snapshots")
            credentials {
                username = System.getenv("MAVEN_USERNAME")
                password = System.getenv("MAVEN_SECRET")
            }
        }
        maven {
            name = "release"
            url = uri("https://repo.emortal.dev/releases")
            credentials {
                username = System.getenv("MAVEN_USERNAME")
                password = System.getenv("MAVEN_SECRET")
            }
        }
    }

    publications {
        create<MavenPublication>("maven") {
            groupId = "dev.emortal.api"
            artifactId = "common-proto-sdk"

            val commitHash = System.getenv("COMMIT_HASH_SHORT")
            val releaseVersion = System.getenv("RELEASE_VERSION")
            version = commitHash ?: releaseVersion ?: "local"

            from(components["java"])
        }
    }
}

// Inform IDEs like IntelliJ IDEA, Eclipse or NetBeans about the generated code.
sourceSets {
    main {
        java {
            srcDirs("build/generated/source/proto/main/grpc")
            srcDirs("build/generated/source/proto/main/java")
        }

        proto {
            srcDir("src/proto")
        }
    }
}
