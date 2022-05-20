package main

import (
	"log"

	pb "blog/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var serverAddress string = "0.0.0.0:50054"

func main() {

	connection, err := grpc.Dial(serverAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Error on connect with the server: %v\n", err)
	}

	defer connection.Close()

	client := pb.NewBlogServiceClient(connection)
	id := createBlog(client)
	blog := readBlog(client, id)
	blog.Title = "Craudios modificou"
	updateBlog(client, blog)
}
