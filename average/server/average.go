package main

import (
	pb "average/averagepb"
	"io"
	"log"
)

func (server *Server) Average(stream pb.AverageService_AverageServer) error {
	log.Println("Average function was invoked")
	res := 0.0
	count := 0
	for {
		req, err := stream.Recv()

		log.Printf("Received: %v\n", req)

		if err == io.EOF {
			return stream.SendAndClose(&pb.AverageResponse{
				Average: res / float64(count),
			})
		}

		if err != nil {
			log.Fatalf("Error on receive request: %v\n", err)
		}

		count++
		res += float64(req.Number)

	}
}
