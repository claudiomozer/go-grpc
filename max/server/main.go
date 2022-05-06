package main

import (
	"log"
	pb "max/maxpb"
	"net"

	"google.golang.org/grpc"
)

var serverAddress string = "0.0.0.0:50055"

type Server struct {
	pb.MaxServiceServer
}

func main() {
	lis, error := net.Listen("tcp", serverAddress)

	if error != nil {
		log.Fatalf("Failed to listen on %v\n", error)
	}

	log.Printf("Listening on %s\n", serverAddress)
	s := grpc.NewServer()
	pb.RegisterMaxServiceServer(s, &Server{})

	if error = s.Serve(lis); error != nil {
		log.Fatalf("Failed to serve on %v\n", error)
	}
}
