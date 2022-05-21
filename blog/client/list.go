package main

import (
	pb "blog/proto"
	"context"
	"io"
	"log"
)

func listBlogs(client pb.BlogServiceClient) {
	log.Println("listBlogs was invoked")

	stream, err := client.ListBlogs(context.Background(), &pb.Empty{})

	if err != nil {
		log.Fatalf("Error while caling ListBlogs: %v\n", err)
	}

	for {
		response, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("An error ocurred on receiving blogs: %v\n", err)
		}

		log.Println(response)
	}

}
