package main

import (
	"context"
	"log"
	pb "sqrt/sqrtpb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doSqrt(client pb.SqrtServiceClient, n int32) {
	log.Println("doSqrt was invoked")
	res, err := client.Sqrt(context.Background(), &pb.SqrtRequest{
		Number: n,
	})

	if err != nil {
		e, ok := status.FromError(err)

		if ok {
			log.Printf("Error message from server %s\n", e.Message())
			log.Printf("Error code from server %s\n", e.Code())

			if e.Code() == codes.InvalidArgument {
				log.Printf("We probaly sent a negative number")
				return
			}
		} else {
			log.Fatalf("Undefined error: %v\n", err)
		}
	}

	log.Printf("Sqrt result %f\n", res.Result)
}
