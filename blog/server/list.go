package main

import (
	pb "blog/proto"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) ListBlogs(in *pb.Empty, stream pb.BlogService_ListBlogsServer) error {
	log.Printf("ListBlogs was invoked with: %v\n", in)
	cur, err := collection.Find(context.Background(), primitive.D{{}})

	if err != nil {
		return status.Errorf(
			codes.NotFound,
			"Cannot find blog with the ID provided",
		)
	}
	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		data := &BlogItem{}
		err := cur.Decode(data)

		if err != nil {
			return status.Errorf(
				codes.Internal,
				fmt.Sprintf("Error while decoding MonogoDB: %v\n", err),
			)
		}

		stream.Send(documentToBlog(data))
	}

	if err = cur.Err(); err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Error reading cursor: %v\n", err),
		)
	}

	return nil
}
