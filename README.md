# grpc-api-specs

This repository contains the proto specifications for gRPC and messaging APIs of EmortalMC.
Some protos (those in a `messages.proto` file) are for use with RabbitMQ and not a gRPC API.

## Contributing

After making any changes to proto definitions, please remember to run the `run/generate.sh` file.
Java generation will automatically run on push, but you will need to run the other languages via this file.