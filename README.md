# gRPC Examples

Scratch Repo for Exploring gRPC/Protobuf

## Generating gRPC from Protobuf definition file

The `games.proto` files is where all of the protobuf definitions live for this project.

### Generating Go Code

As long as you have the general `protoc` compiler installed as well as the Go-specific compiler, you can run the following command to generate the code

```sh
protoc -I=./ --go_out=./ ./games.proto
```
