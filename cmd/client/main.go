// This package contains the client for the gRPC service.
package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "grpc-example/pkg/grpc-games"
)

const serverAddr = "localhost:8080"

func main() {

	// Set up the Client
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.NewClient(serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewGameServiceClient(conn)

	// Make requests to the server
	ctx := context.Background()
	// Get a game
	g, err := client.GetGame(ctx, &pb.GetGameRequest{Name: "Zelda"})
	if err != nil {
		log.Fatalf("could not get game: %v", err)
	}
	fmt.Printf("Got Game: \n- %v\n", g)

	// Create a game
	newGame := &pb.Game{Name: "Mario", Genre: pb.Genre_ADVENTURE}
	g, err = client.CreateGame(ctx, newGame)
	if err != nil {
		log.Fatalf("could not create game: %v", err)
	}
	fmt.Printf("Created Game: \n- %v\n", g)

	// Get the created game
	g, err = client.GetGame(ctx, &pb.GetGameRequest{Name: "Mario"})
	if err != nil {
		log.Fatalf("could not get game: %v", err)
	}
	fmt.Printf("Got Game: \n- %v\n", g)
}
