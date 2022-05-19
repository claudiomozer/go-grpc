package main

import (
	pb "blog/proto"
	"context"
	"log"
)

func readBlog(client pb.BlogServiceClient, id string) *pb.Blog {
	log.Println("readBlog was invoked")

	req := &pb.BlogId{Id: id}

	res, err := client.ReadBlog(context.Background(), req)

	if err != nil {
		log.Printf("Error happened while reading: %v\n", err)
	}

	log.Printf("Blog was read: %v\n", res)
	return res
}
