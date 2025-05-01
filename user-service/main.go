package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "proto"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedUserServiceServer
}

var users = map[int32]string{
	1: "Alice",
	2: "Bob",
}

func (s *server) GetUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	name, exists := users[req.Id]
	if !exists {
		return nil, fmt.Errorf("User not found")
	}
	return &pb.UserResponse{Id: req.Id, Name: name}, nil
}

func main() {
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &server{})

	fmt.Println("User service running on gRPC port 50051...")
	if err := s.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
