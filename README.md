# proto-specs

This repository contains the proto specifications for gRPC and messaging APIs of EmortalMC.
Some protos (those in a `messages.proto` file) are for use with RabbitMQ and not a gRPC API.

## Contributing

After making any changes to proto definitions, please remember to run the `run/generate.sh` file.
Java generation will automatically run on push, but you will need to run the other languages via this file.

## Service Ports

Note, these ports are not used in production, but are defaults for local development.
In production, the ports of default gRPC services are locked to 80.

Any services that share a port mean they are combined within the same codebase
and are run on the same gRPC server.

| Service                    | Description                                         | Port  |
|----------------------------|-----------------------------------------------------|-------|
| permission                 | Manages permissions                                 | 10001 |
| relationship               | Manages friends & blocks                            | 10002 |
| message-handler            | Handles private messages + chat messages            | 10003 |
| mc-player                  | Manages core user data (username, playtime, etc.)   | 10004 |
| badge-manager              | Manages badges                                      | 10004 |
| player-tracker             | Tracks player counts & player proxy/server          | 10005 |
| party-manager              | Manages parties                                     | 10006 |
| matchmaker (kurushimi)     | Matchmaker, separate repo for gRPC spec             | 10007 |
| account-connection-manager | Manages account connections (Mojang, Discord, etc.) | 10008 |
