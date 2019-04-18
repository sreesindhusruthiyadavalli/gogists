package main

import (
	"log"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	//pb "google.golang.org/grpc/examples/helloworld/helloworld"
)

const (
	address = "localhost:50051"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := NewMyMathServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Add(ctx, &NumbersRequest{A: 3, B: 5})
	if err != nil {
		log.Fatalf("Error calling grpc Add: %v", err)
	}
	log.Printf("Sum of values is: %d", r.Ans)
}
