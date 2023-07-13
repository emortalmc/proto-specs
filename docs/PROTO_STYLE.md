# Protobuf Style Guide

This guide extends from the official style guide, in that everything outlined in the official style guide is also
recommended to be used here. As such, anything included in the official style guide has been omitted here for the
sake of brevity.

The official style guide can be found [here](https://protobuf.dev/programming-guides/style/).

Every file in this repository should use the proto version 3 syntax.

## Service naming

* Every service should be named in PascalCase.
* The use of generic terms such as 'Service' or 'Manager' in the service name should be avoided (e.g. call it `McPlayer` not `McPlayerService`)

## Directory structure

* Every service should have its own directory where only its declarations live.
* This directory should be named the same as the service but in snake_cast (e.g. `mc_player`)

## File naming

Each service is generally made up of three components:
* service RPC declarations, requests/responses and error responses (`grpc.proto`)
* models that are used by the service to model data (`models.proto`), e.g. a user, player, or game type
* the messages (events) that can be sent by the service over Kafka (`messages.proto`), e.g. a user being created, or a game being started

## Package naming

### Protobuf
* The full name should be `emortal.<component>.<service>` where the component is either `grpc`, `model`, or `message`
* The service part of the package name should be in lower snake_case

An example of this would be for a service called McPlayer the package will be `emortal.grpc.mc_player`

### Java
* The full name should be `dev.emortal.api.<component>.<service>` where the component is either `grpc`, `model`, or `message`
* The service part of the package name should be in lower case with no spaces

An example of this would be for a service called McPlayer the package will be `dev.emortal.api.grpc.mcplayer`

### Go
* The full name should be `github.com/emortalmc/proto-specs/gen/go/<component>/<service>`
* The service part of the package name should be in lower case with no spaces (same as Java)

An example of this would be for a service called McPlayer the package will be `github.com/emortalmc/proto-specs/gen/go/grpc/mcplayer`

## Documentation

All elements should be well documented, as advised in the [API Best Practices](https://protobuf.dev/programming-guides/api/).

The general rule of thumb is try to document things like they are going to be read by someone who has no idea how
any of the system works, and is trying to figure it out.

## Generated Code Options

### Java

TODO - Apply this convention everywhere

For gRPC declarations (`grpc.proto` file), the option `java_outer_classname` should be set to the name of the service (in pascal case),
with the suffix "Proto" appended to it. For example, the name for a service called "McPlayer" would be "McPlayerProto".

For other declarations (`messages.proto` and `models.proto`), the option `java_multiple_files` should be used instead
which will generate a separate Java file for each message declaration.

## Errors

Always prefer using existing status codes over a custom error response, however, don't stretch it to a point of
ambiguity and don't use a combination of both status codes and custom error responses. Use custom error codes when:
* There are multiple errors that fit into the same status code
* There isn't a clear status code that fits the error
* You want to provide additional information or context to the error to the client

If existing status codes aren't sufficient you should define your own error response named `<RequestName>ErrorResponse`
(e.g. `CreateUserErrorResponse`) and include it in the service's `grpc.proto` file.
It should be after the `<RequestName>Response` message. This message should contain an enum called `Reason` and be
specified in a field called `reason` at index 1.