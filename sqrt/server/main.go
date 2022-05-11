package main

import (
	"log"
	"net"
	pb "sqrt/sqrtpb"

	"google.golang.org/grpc"
)

var serverAddress string = "0.0.0.0:50052"

type Server struct {
	pb.SqrtServiceServer
}

func main() {
	tcpListener, err := net.Listen("tcp", serverAddress)

	if err != nil {
		log.Fatal("Error on open TCP server")
	}
	log.Printf("Server listening on %s\n", serverAddress)

	grpcServer := grpc.NewServer()
	pb.RegisterSqrtServiceServer(grpcServer, &Server{})

	err = grpcServer.Serve(tcpListener)
	if err != nil {
		log.Fatalf("Error on serve gRPC server: %v\n", err)
	}
}
