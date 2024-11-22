// This package contains the server for the gRPC service.
package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "grpc-example/pkg/grpc-games"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

type server struct {
	pb.UnimplementedGameServiceServer
	savedGames []*pb.Game
}

func newServer() *server {
	s := &server{
		savedGames: []*pb.Game{},
	}
	return s
}

func (s *server) GetGame(ctx context.Context, req *pb.GetGameRequest) (*pb.Game, error) {
	for _, game := range s.savedGames {
		if game.Name == req.Name {
			return game, nil
		}
	}
	return nil, status.Errorf(codes.NotFound, "game not found")
}

func (s *server) CreateGame(ctx context.Context, req *pb.Game) (*pb.Game, error) {
	s.savedGames = append(s.savedGames, req)
	return req, nil
}

func main() {
	port := 8080
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	fmt.Println("Server listening on port", port)

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterGameServiceServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}
