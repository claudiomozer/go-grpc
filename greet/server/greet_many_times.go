package main

import (
	"fmt"
	pb "greet/greetpb"
	"log"
)

func (server *Server) GreetManyTimes(request *pb.GreetRequest, stream pb.GreetService_GreetManyTimesServer) error {
	log.Printf("GreetManyTimes function was invoked with: %v\n", request)

	for i := 0; i < 10; i++ {
		res := fmt.Sprintf("Hello %s, number %d", request.FirstName, i)

		stream.Send(&pb.GreetResponse{
			Result: res,
		})
	}
	return nil
}
