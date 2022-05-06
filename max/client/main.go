package main

import (
	"log"
	pb "max/maxpb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var serverAddress string = "0.0.0.0:50055"

func main() {
	conn, err := grpc.Dial(serverAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}

	defer conn.Close()

	client := pb.NewMaxServiceClient(conn)
	doMax(client)
}
