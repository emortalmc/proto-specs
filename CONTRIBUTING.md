# How to Contribute

We'd love to accept your contributions to this project - there are just a few small guidelines you need to follow.

## Submitting code via pull requests
* All code must be submitted via pull requests.
* We follow the [GitHub Pull Request Model](https://docs.github.com/en/pull-requests/collaborating-with-pull-requests/proposing-changes-to-your-work-with-pull-requests/about-pull-requests) for all contrubitions.
* For large bodies of work, please open an issue first to discuss the work with the team describing the feature you
would like to build and how you plan to implement it.
* All submissions will require review before being merged.
* Please follow our style guides for [**Java**](todo), [**Golang**](todo) and [**Protobuf**](docs/PROTO_STYLE.md).

## Additional instructions for protobuf changes
Make sure to run the `gen/generate.sh` script before committing any changes. This will ensure changes to protobuf
files are reflected in the generated code. Java code is auto-generated on gradle build.

## Project Layout

Protobuf files are located in the `src/proto` directory.