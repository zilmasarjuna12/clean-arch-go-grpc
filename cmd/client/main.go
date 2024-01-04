package main

import (
	product_grpc "clean-arch-go-grpc/internal/delivery/grpc/proto"
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:3000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := product_grpc.NewProductHandlerClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.GetList(ctx, &product_grpc.Empty{})

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("ID: %s", r)
}
