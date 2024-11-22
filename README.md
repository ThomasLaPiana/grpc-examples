# gRPC Examples

Scratch Repo for Exploring gRPC/Protobuf

The key features of Protobufs/gRPC:

1. Use of `.proto` files to declare the messages (request/response objects)
2. Automatically generated client and server code
3. High-performance (compared to REST)
4. Type information encoded as part of the message

[More info explaining REST vs gRPC and when to consider each](https://aws.amazon.com/compare/the-difference-between-grpc-and-rest/)

## Generating gRPC from Protobuf definition file

The `games.proto` files is where all of the `Protobuf` and `gRPC` definitions live for this project.

### Generating Client & Server Code

As long as you have the general `protoc` compiler installed as well as the Go-specific compilers (protobuf as well as gRPC), you can run the following command to generate the code

```sh
protoc --go_out=./pkg/grpc-games --go_opt=paths=source_relative --go-grpc_out=./pkg/grpc-games --go-grpc_opt=paths=source_relative games.proto
```

Once written, using the command below generates two different files:

* `pkg/grpc-games/games.pb.go` - This file contains the generated Go types
* `pkg/grpc-games/games_grpc.pb.go` - This file contains the generated service endpoints

While `gRPC` is responsible for generating the client & interface definition, on the server-side it is still up to us to fulfill the contract. This requires creating a server and application that fulfills the interface.

## Running the Client & Server

1. In one terminal, run `go run ./cmd/server`
2. In another terminal, run `go run ./cmd/client`
