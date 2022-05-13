package main

import (
	pb "greet/greetpb"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var serverAddress string = "localhost:50051"

type Server struct {
	pb.GreetServiceServer
}

func main() {
	lis, error := net.Listen("tcp", serverAddress)

	if error != nil {
		log.Fatalf("Failed to listen on %v\n", error)
	}

	log.Printf("Listening on %s\n", serverAddress)

	tls := true
	opts := []grpc.ServerOption{}

	if tls {
		certFile := "/home/claudio/workspace/go/src/go-course/ssl/server.crt"
		keyFile := "/home/claudio/workspace/go/src/go-course/ssl/server.pem"
		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)

		if err != nil {
			log.Fatalf("Error on loading certificates: %v\n", err)
		}

		opts = append(opts, grpc.Creds(creds))
	}

	s := grpc.NewServer(opts...) // ... it's the spread operator on golang
	pb.RegisterGreetServiceServer(s, &Server{})

	if error = s.Serve(lis); error != nil {
		log.Fatalf("Failed to serve on %v\n", error)
	}
}
