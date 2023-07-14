package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/vergissberlin/go-tutorial-grpc/greet/proto"
)

var addr string = "0.0.0.0:50012"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	defer conn.Close()
	c := pb.NewGreetServiceClient(conn)

	// Call the doGreat function
	doGreet(c)

}
