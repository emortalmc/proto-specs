# Protobuf specification format

This file is intended to provide a guide on how to name and structure Protobuf declarations in this repository.

It is designed to make sure that all our declarations are consistent, and that the practices are well documented,
well understood, and easy to follow, so that everyone stays on the same page as to how things are meant to work,
and the designs are not completely inconsistent.

This guide extends from the official style guide, in that everything outlined in the official style guide is also
recommended to be used here. As such, anything included in the official style guide has been omitted here for the
sake of brevity.

The official style guide can be found [here](https://protobuf.dev/programming-guides/style/).

Every file in this repository should use the proto version 3 syntax.

# ⚠️ TODOs ⚠️

## Service naming

Every service should be named in PascalCase. The use of generic terms such as 'Service' or 'Manager' in the
service name should be avoided.

## Directory structure

Every service should have its own directory, where only its declarations live. This directory should be named the
same as the service, so the declarations are easy to find when searching through the repository.

## File naming

Each service is generally made up of three components: the service RPC declarations, the models that are used by the
service to model data, and the messages (events) that can be sent by the service over Kafka.

These three components should be separated into their own files, and named as follows:
* RPC declarations: `grpc.proto`. This file should contain all service declarations, and any messages required for
  the service, such as requests, responses, and error responses.
* Models: `models.proto`. This file should contain any data models used by the service, such as a user, player, or
  game type.
* Messages: `messages.proto`. This file should contain any Kafka messages that may be sent by the service to indicate
  that something has happened, such as a user being created, or a game being started.

## Package naming

Every package name in every proto file should be prefixed with the namespace `emortal`. This is for two reasons:
1. The namespace is unique, and so there is no chance of conflict with dependent packages.
2. It ensures that there is a consistent place to find all generated code. In Java, for example, all the code will
   be located in the `dev.emortal.api` package, under various consistent sub-packages listed below.

In addition, the package naming should follow this format: `emortal.<component>.<service>`, where the component is
either `grpc`, `model`, or `message`, and the service is the name of the service that the declarations are for.

For example, the gRPC declarations for a service called `user` would be in the package `emortal.grpc.user`.

The service part of the package name should be in lower snake case, for the Protobuf package. Generated packages
are different, and are covered below.

### Generated code

The package naming for the generated code is different depending on the language. The format is as follows
for each language:
* Java: `dev.emortal.api.<component>.<service>`
* Go: `github.com/emortalmc/proto-specs/gen/go/<component>/<service>`

The service part of the generated package name should be as follows for each language:
* Java: concatenate each word in the service name together, rather than separating them with underscores, as is
  recommended by the Google Java style guide, which is recommended for usage in
  the [Proto Best Practices](https://protobuf.dev/programming-guides/dos-donts/).
* Go: concatenate each word in the service name together, rather than separating them with underscores, as is
  recommended by [Effective Go](https://go.dev/doc/effective_go#package-names).

## Documentation

All elements should be well documented, as advised in the [API Best Practices](https://protobuf.dev/programming-guides/api/).

The general rule of thumb is try to document things like they are going to be read by someone who has no idea how
any of the system works, and is trying to figure it out.

Every RPC method should have a description of the expected behaviour, and should document any errors that may occur
as part of expected behaviour, for example, invalid parameters, or a user not being found.

## Generated code options

### Java

TODO - apply this naming convention everywhere

For gRPC declarations, the option `java_outer_classname` should be set to the name of the service, in pascal case,
with the suffix "Proto" appended to it. For example, the name for a service called "User"
would be "UserProto".

For other declarations, the option `java_multiple_files` should be used instead, which will generate a separate
Java file for each top-level declaration.

## Errors

Always prefer to use a pre-existing status code over a custom error type. Not only are they more consistent and
widely known, they are also a lot simpler to handle for client code, as they only have to check the status.

If none of the pre-existing status codes fit the error that is being returned, then a custom error response should
be defined, containing an enum type with error status codes, along with any data that may be useful for clients
to debug the error.
