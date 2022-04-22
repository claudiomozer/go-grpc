package main

import (
	"context"
	pb "greet/greetpb"
	"log"
)

func doGreet(client pb.GreetServiceClient) {
	log.Println("doGreet was invoked")
	res, err := client.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Claufio",
	})

	if err != nil {
		log.Fatalf("Could not invoke Greet: %v\n", err)
	}

	log.Printf("Greeting %s\n", res.Result)
}
