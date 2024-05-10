# Ecumenos protobuf

This folder contains interface definition of all services.
All generated `.go` code is under `gen/` folder.

For protobuf file, we may want to follow these:

- all APIs support both REST and gRPC protocols;
- follow Google's API guide, use `Standard Methods` as much as it is possible for us;
- provide comments for all APIs;

## Requirements

- golang >= 1.21.5

## Usage
