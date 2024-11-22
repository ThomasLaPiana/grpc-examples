# gRPC Examples

Scratch Repo for Exploring gRPC/Protobuf

## Generating gRPC from Protobuf definition file

The `games.proto` files is where all of the protobuf definitions live for this project.

Once written, using the command below generates two different files:

* `games.pb.go` - This file contains the generated Go types
* `games_grpc.pb.go` - This file contains the generated service endpoints

### Generating Go Code

As long as you have the general `protoc` compiler installed as well as the Go-specific compilers (protobuf as well as gRPC), you can run the following command to generate the code

```sh
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    games.proto
```
