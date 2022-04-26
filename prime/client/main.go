package main

import (
	"log"
	pb "prime/primepb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var serverAddress string = "0.0.0.0:50053"

func main() {
	conn, err := grpc.Dial(serverAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}

	defer conn.Close()

	client := pb.NewPrimeServiceClient(conn)

	decompose(client)
}
