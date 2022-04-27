package main

import (
	"context"
	pb "greet/greetpb"
	"log"
	"time"
)

func doLongGreet(client pb.GreetServiceClient) {
	log.Println("doLongGreet was invoked")

	reqs := []*pb.GreetRequest{
		{FirstName: "Claudio"},
		{FirstName: "Craudio"},
		{FirstName: "Claudinho"},
	}

	stream, err := client.LongGreet(context.Background())

	if err != nil {
		log.Fatalf("Error on send requests %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while receiving response from LongGreet: %v\n", err)
	}

	log.Printf("LongGreet, %s\n", res.Result)
}
