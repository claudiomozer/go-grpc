package main

import (
	pb "greet/greetpb"
	"log"
	"net"

	"google.golang.org/grpc"
)

var serverAddress string = "0.0.0.0:50051"

type Server struct {
	pb.GreetServiceServer
}

func main() {
	lis, error := net.Listen("tcp", serverAddress)

	if error != nil {
		log.Fatalf("Failed to listen on %v\n", error)
	}

	log.Printf("Listening on %s\n", serverAddress)
	s := grpc.NewServer()
	pb.RegisterGreetServiceServer(s, &Server{})

	if error = s.Serve(lis); error != nil {
		log.Fatalf("Failed to serve on %v\n", error)
	}
}
