package main

import (
	"context"
	"log"

	pb "github.com/vergissberlin/go-tutorial-grpc/greet/proto"
)

func doGreat(client pb.GreetServiceClient) {
	log.Println("doGreat() was invoked")
	res, err := client.Greet(context.Background(), &pb.GreetingRequest{
		FirstName: "Vergiss",
		LastName:  "Berlin",
	})

	if err != nil {
		log.Fatalf("Error while calling Greet RPC: %v\n", err)
		panic(err)
	}

	log.Printf("Response from Greet: %s\n", res.Result)

}
