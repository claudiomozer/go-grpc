package main

import (
	"context"
	"io"
	"log"
	pb "max/maxpb"
	"time"
)

func doMax(client pb.MaxServiceClient) {
	log.Printf("doMax was invoked")

	stream, err := client.Max(context.Background())

	if err != nil {
		log.Fatalf("Error while creating stream: %v\n", err)
	}

	reqs := []*pb.MaxRequest{
		{Number: -1},
		{Number: -2},
		{Number: 0},
		{Number: 7},
		{Number: 6},
		{Number: 4},
		{Number: 1},
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

			log.Printf("Max value %d\n", res.Result)
		}
		close(waitc)
	}()
	<-waitc
}
