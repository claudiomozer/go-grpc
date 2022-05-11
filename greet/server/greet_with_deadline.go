package main

import (
	"context"
	pb "greet/greetpb"
	"log"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GreetWithDeadline(ctx context.Context, request *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greet with deadline was invoked with %v\n", request)
	for i := 0; i < 3; i++ {
		if ctx.Err() == context.DeadlineExceeded {
			log.Println("The client canceled the request!")
			return nil, status.Error(codes.Canceled, "The client canceled the request")
		}
		time.Sleep((1 * time.Second))
	}

	return &pb.GreetResponse{
		Result: "Hello " + request.FirstName,
	}, nil
}
