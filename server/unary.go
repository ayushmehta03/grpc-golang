package main

import (
	"context"

	pb "github.com/ayushmehta03/grpc-golang/proto"
)

//unary api-> single req and response model

func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Hello",
	}, nil
}