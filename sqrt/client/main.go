package main

import (
	"log"
	pb "sqrt/sqrtpb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var serverAddress string = "0.0.0.0:50052"

func main() {
	conn, err := grpc.Dial(serverAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}

	defer conn.Close()
	serviceClient := pb.NewSqrtServiceClient(conn)
	doSqrt(serviceClient, -10)
}
