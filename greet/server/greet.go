package main

import (
	"context"
	"log"

	pb "github.com/vergissberlin/go-tutorial-grpc/greet/proto"
)

func (*Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greet was invoked with %v\n", in)
	return &pb.GreetResponse{Result: "Hello " + in.FirstName + " " + in.LastName + "!"}, nil
}
