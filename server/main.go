package main

import (
	"net"

	parkingpb "parky/proto"
	"parky/server/services"

	"github.com/golang/glog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		glog.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	parkingpb.RegisterAuthenticationServer(s, &services.AuthenticationServer{})

	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		glog.Fatalf("failed to serve: %v", err)
	}
}
