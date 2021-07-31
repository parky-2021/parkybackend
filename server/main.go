package main

import (
	"log"
	"net"

	"context"
	parkingpb "parky/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	parkingpb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *parkingpb.HelloRequest) (*parkingpb.HelloReply, error) {
	return &parkingpb.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	parkingpb.RegisterGreeterServer(s, &server{})

	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
