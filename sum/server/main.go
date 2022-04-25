package main

import (
	"log"
	"net"
	pb "sum/sumpb"

	"google.golang.org/grpc"
)

var serverAddress string = "0.0.0.0:50052"

type Server struct {
	pb.SumServiceServer
}

func main() {
	tcpListener, err := net.Listen("tcp", serverAddress)

	if err != nil {
		log.Fatal("Error on open TCP server")
	}
	log.Printf("Server listening on %s\n", serverAddress)

	grpcServer := grpc.NewServer()
	pb.RegisterSumServiceServer(grpcServer, &Server{})

	err = grpcServer.Serve(tcpListener)
	if err != nil {
		log.Fatalf("Error on serve gRPC server: %v\n", err)
	}
}
