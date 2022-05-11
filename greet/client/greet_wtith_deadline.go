package main

import (
	"context"
	pb "greet/greetpb"
	"log"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doGreetWithDeadline(client pb.GreetServiceClient, timeout time.Duration) {
	log.Println("doGreetWithDeadline was invoked")

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	request := &pb.GreetRequest{
		FirstName: "Claudio",
	}

	res, err := client.GreetWithDeadline(ctx, request)

	if err != nil {
		e, ok := status.FromError(err)

		if ok {
			if e.Code() == codes.DeadlineExceeded {
				log.Println("Deadline exceeded")
				return
			} else {
				log.Fatalf("Unexpected grpc error from server: %v\n", err)
			}
		} else {
			log.Fatalf("Unexpected error from server: %v\n", err)
		}

	}

	log.Printf("GreetWithDeadline %s\n", res.Result)
}
