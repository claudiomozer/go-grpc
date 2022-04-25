package main

import (
	"context"
	"log"
	pb "sum/sumpb"
)

func sum(client pb.SumServiceClient) {
	res, err := client.Sum(context.Background(), &pb.SumRequest{
		FirstValue:  13,
		SecondValue: 7,
	})

	if err != nil {
		log.Fatalf("Error on sum values: %v\n", err)
	}
	log.Printf("Sum result %d\n", res.Result)
}
