package main

import (
	"context"
	"log"
	pb "sum/sumpb"
)

func (server *Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("sum was invoked %v\n", in)
	firstValue := in.FirstValue
	secondValue := in.SecondValue
	var sumResult int32 = firstValue + secondValue

	return &pb.SumResponse{
		Result: sumResult,
	}, nil
}
