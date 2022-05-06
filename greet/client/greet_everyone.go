package main

import (
	"context"
	pb "greet/greetpb"
	"io"
	"log"
	"time"
)

func doGreetEveryone(client pb.GreetServiceClient) {
	log.Printf("doGreetEveryone was invoked")

	stream, err := client.GreetEveryone(context.Background())

	if err != nil {
		log.Fatalf("Error while creating stream: %v\n", err)
	}

	reqs := []*pb.GreetRequest{
		{FirstName: "Claudio"},
		{FirstName: "Craudio"},
		{FirstName: "Claudin"},
		{FirstName: "Craudin"},
		{FirstName: "Claudinei"},
	}

	waitc := make(chan struct{})

	go func() {
		for _, req := range reqs {
			log.Printf("Sending message %v\n", req)
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Printf("Error on receive response %v\n", err)
				break
			}

			log.Printf("Received %v\n", res.Result)
		}
		close(waitc)
	}()
	<-waitc
}
