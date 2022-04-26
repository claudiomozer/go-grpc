package main

import (
	"context"
	"io"
	"log"
	pb "prime/primepb"
)

func decompose(client pb.PrimeServiceClient) {
	log.Println("decompose was invoked")

	req := &pb.PrimeRequest{
		Number: 56,
	}

	stream, err := client.Prime(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while calling decompose: %v\n", err)
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading the stream: %v\n", err)
		}

		log.Printf("Decomposition: %d\n", msg.Decomposition)
	}
}
