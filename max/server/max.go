package main

import (
	"io"
	"log"
	pb "max/maxpb"
)

func (s *Server) Max(stream pb.MaxService_MaxServer) error {
	log.Printf("Max was invoked")
	var max *int64

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error on receiving request: %v\n", err)
		}

		if max == nil {
			max = &req.Number
		}

		if req.Number > *max {
			*max = req.Number
		}

		err = stream.Send(&pb.MaxResponse{
			Result: *max,
		})

		if err != nil {
			log.Fatalf("Error on sending request: %v\n", err)
		}
	}
}
