package main

import (
	pb "blog/proto"
	"context"
	"log"
)

func deleteBlog(client pb.BlogServiceClient, id string) {
	log.Println("deleteBlog was invoked")

	req := &pb.BlogId{Id: id}

	res, err := client.DeleteBlog(context.Background(), req)

	if err != nil {
		log.Printf("Error happened while deleting: %v\n", err)
	}

	log.Printf("Blog was read: %v\n", res)
}
