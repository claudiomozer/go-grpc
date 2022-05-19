package main

import (
	pb "blog/proto"
	"context"
	"log"
)

func createBlog(client pb.BlogServiceClient) string {
	log.Println("--createBlog was invoked--")

	blog := &pb.Blog{
		AuthorId: "teste",
		Title:    "Meu primeiro blog",
		Content:  "Conteúdo fictício do meu primeiro blog",
	}
	res, err := client.CreateBlog(context.Background(), blog)

	if err != nil {
		log.Fatalf("Unexpected error on creating blog: %v\n", err)
	}

	log.Printf("Blog hase been created: %s\n", res.Id)
	return res.Id
}
