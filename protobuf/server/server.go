package main

import (
	"fmt"
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

// server is used to implement the service
type server struct{}

func (s *server) Add(ctx context.Context, in *NumbersRequest) (*AnswerReply, error) {
	ans := in.A + in.B
	hr := AnswerReply{
		Ans: ans,
	}
	return &hr, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	fmt.Println("Listening at ", port)
	s := grpc.NewServer()
	RegisterMyMathServiceServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
