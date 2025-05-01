package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/HarshPanchal18/Go-Micro/proto"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := client.GetUser(ctx, &pb.UserRequest{Id: 1})
	if err != nil {
		log.Fatalf("could not get user: %v", err)
	}

	fmt.Printf("User: ID=%d, Name=%s\n", resp.Id, resp.Name)
}
