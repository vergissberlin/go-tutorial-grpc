package main

import (
	"context"
	"log"

	pb "github.com/vergissberlin/go-tutorial-grpc/greet/proto"
)

func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greet function was invoked with %v\n", in)
	return &pb.GreetResponse{
		Result: "Hello " + in.FirstName + " " + in.LastName + "!",
	}, nil
}
