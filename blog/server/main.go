package main

import (
	pb "blog/proto"
	"context"
	"log"
	"net"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

var collection *mongo.Collection
var serverAddress string = "0.0.0.0:50054"

type Server struct {
	pb.BlogServiceServer
}

func main() {

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:root@localhost:27017/"))

	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(context.Background())

	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database("blogdb").Collection("blog")

	listener, error := net.Listen("tcp", serverAddress)

	if error != nil {
		log.Fatalf("Error on listen to address: %v\n", error)
	}

	log.Printf("Server listening on %s\n", serverAddress)

	grpcServer := grpc.NewServer()
	pb.RegisterBlogServiceServer(grpcServer, &Server{})

	error = grpcServer.Serve(listener)

	if error != nil {
		log.Fatalf("Error on grpc serve %v\n", error)
	}
}
