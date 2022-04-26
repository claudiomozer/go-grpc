package main

import (
	"log"
	pb "prime/primepb"
)

func (server *Server) Prime(request *pb.PrimeRequest, stream pb.PrimeService_PrimeServer) error {
	log.Printf("Prime decomposition was invoked with: %v\n", request)

	var valueToDecompose int32 = request.Number
	var divisor int32 = 2
	for {

		if (valueToDecompose % divisor) == 0 {
			valueToDecompose = valueToDecompose / divisor
			stream.Send(&pb.PrimeResponse{
				Decomposition: divisor,
			})
		} else {
			divisor += 1
		}

		if valueToDecompose == 1 {
			break
		}
	}
	return nil
}
