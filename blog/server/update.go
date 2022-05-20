package main

import (
	pb "blog/proto"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) UpdateBlog(ctx context.Context, in *pb.Blog) (*pb.Empty, error) {
	log.Printf("UpdateBlog was invoked with: %v\n", in)

	oid, err := primitive.ObjectIDFromHex(in.Id)

	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Cannot parse ID",
		)
	}

	blogItem := &BlogItem{
		AuthorId: in.AuthorId,
		Content:  in.Content,
		Title:    in.Title,
	}
	data := bson.M{"$set": blogItem}
	filter := bson.M{"_id": oid}

	res, err := collection.UpdateOne(ctx, filter, data)

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Error on update: %v\n", err),
		)
	}

	if res.MatchedCount == 0 {
		return nil, status.Errorf(
			codes.NotFound,
			"Cannot find blog with the ID provided",
		)
	}

	return &pb.Empty{}, nil
}
