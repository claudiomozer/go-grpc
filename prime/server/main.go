package main

import (
	"log"
	"net"
	pb "prime/primepb"

	"google.golang.org/grpc"
)

var serverAddress string = "0.0.0.0:50053"

type Server struct {
	pb.PrimeServiceServer
}

func main() {
	lis, error := net.Listen("tcp", serverAddress)

	if error != nil {
		log.Fatalf("Failed to listen on %v\n", error)
	}

	log.Printf("Listening on %s\n", serverAddress)
	s := grpc.NewServer()
	pb.RegisterPrimeServiceServer(s, &Server{})

	if error = s.Serve(lis); error != nil {
		log.Fatalf("Failed to serve on %v\n", error)
	}
}
