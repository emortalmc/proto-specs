import com.google.protobuf.gradle.*

plugins {
    `java-library`
    `maven-publish`

    // gRPC
    id("com.google.protobuf") version "0.9.4"
}

group = "dev.emortal.api"
version = "1.0"

val grpcVersion = "1.63.0"
val protobufVersion = "4.26.1"
val protocVersion = protobufVersion

repositories {
    mavenCentral()

    maven("https://packages.confluent.io/maven/")
}

dependencies {
    api("com.google.protobuf:protobuf-java:$protobufVersion")
    api("org.apache.tomcat:tomcat-annotations-api:10.1.23")
    api("org.slf4j:slf4j-api:2.0.13")

    implementation("io.micrometer:micrometer-core:1.12.5")
    implementation("org.apache.kafka:kafka-clients:7.6.0-ccs")

    compileOnly("org.jetbrains:annotations:24.1.0")
    annotationProcessor("org.jetbrains:annotations:24.1.0")

    // Protobuf uses @javax.annotation.Generated on generated types, and sometimes this isn't available.
    // This make sure it's always available
    compileOnly("javax.annotation:javax.annotation-api:1.3.2")

    implementation("com.github.ben-manes.caffeine:caffeine:3.1.8")

    api("io.grpc:grpc-netty:$grpcVersion")
    api("io.grpc:grpc-stub:$grpcVersion")
    api("io.grpc:grpc-protobuf:$grpcVersion")

    testImplementation("org.junit.jupiter:junit-jupiter-api:5.10.2")
    testImplementation("org.junit.jupiter:junit-jupiter-params:5.10.2")
    testRuntimeOnly("org.junit.jupiter:junit-jupiter-engine:5.10.2")
    testRuntimeOnly("org.slf4j:slf4j-simple:2.0.13")
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
        languageVersion.set(JavaLanguageVersion.of(21))
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
