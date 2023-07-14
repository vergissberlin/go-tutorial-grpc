package main

import (
	"context"
	"log"

	pb "github.com/vergissberlin/go-tutorial-grpc/greet/proto"
)

func (s *Server) Greet(ctx context.Context, in *pb.GreetingRequest) (*pb.GreetingResponse, error) {
	log.Printf("Greet function was invoked with %v\n", in)
	return &pb.GreetingResponse{
		Result: "Hello " + in.FirstName + " " + in.LastName + "!",
	}, nil
}
