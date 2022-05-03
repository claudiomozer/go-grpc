package main

import (
	pb "average/averagepb"
	"log"
	"net"

	"google.golang.org/grpc"
)

var serverAddress string = "0.0.0.0:50054"

type Server struct {
	pb.AverageServiceServer
}

func main() {

	listener, error := net.Listen("tcp", serverAddress)

	if error != nil {
		log.Fatalf("Error on listen to address: %v\n", error)
	}

	log.Printf("Server listening on %s\n", serverAddress)

	grpcServer := grpc.NewServer()
	pb.RegisterAverageServiceServer(grpcServer, &Server{})

	error = grpcServer.Serve(listener)

	if error != nil {
		log.Fatalf("Error on grpc serve %v\n", error)
	}
}
