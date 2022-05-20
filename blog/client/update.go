package main

import (
	pb "blog/proto"
	"context"
	"log"
)

func updateBlog(client pb.BlogServiceClient, blog *pb.Blog) *pb.Empty {
	log.Printf("updateBlog was invoked with %v\n", blog)

	_, err := client.UpdateBlog(context.Background(), blog)

	if err != nil {
		log.Fatalf("Error happened while updating: %v\n", err)
	}

	log.Println("Blog was updated")
	return nil
}
