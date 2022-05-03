package main

import (
	pb "average/averagepb"
	"context"
	"log"
	"time"
)

func doAvarage(client pb.AverageServiceClient) {
	log.Println("doAverage was invoked")

	requests := []*pb.AverageRequest{
		{Number: 1},
		{Number: 1},
		{Number: 1},
		{Number: 1},
	}

	stream, err := client.Average(context.Background())

	if err != nil {
		log.Fatalf("Error on get stream: %v\n", err)
	}

	for _, req := range requests {
		log.Printf("Sending request: %v", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error on receive response from server: %v\n", err)
	}

	log.Printf("Average: %f\n", res.Average)
}
