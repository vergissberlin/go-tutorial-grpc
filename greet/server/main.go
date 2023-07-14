package main

import (
	"log"
	"net"

	pb "github.com/vergissberlin/go-tutorial-grpc/greet/proto"
)

var addr string = "0.0.0.0:50012"



// Server represents the gRPC server
type Server struct{
	pb.GreetServiceServer
}

func main() {
	// Listen for incoming gRPC connections.
	list, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	log.Printf("Listening on %s\n", addr)

	// Create a gRPC server object
	srv := grpc.NewServer()
	if err = srv.Serve(list); err != nil {
		panic(err)
	}
	// Register the server as an implementation of the service
	pb.RegisterGreetServiceServer(srv, &Server{})
	
	
}
