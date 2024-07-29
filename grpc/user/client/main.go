package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	pb "user/proto"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		return
	}

	//conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	//if err != nil {
	//	log.Fatalf("Failed to connect: %v", err)
	//}
	defer conn.Close()

	c := pb.NewUserServiceClient(conn)

	//ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	//defer cancel()

	res, err := c.GetUser(context.Background(), &pb.GetUserRequest{Id: "1"})
	if err != nil {
		log.Fatalf("Error calling GetUser: %v", err)
	}

	log.Printf("User: %s, Email: %s", res.GetName(), res.GetEmail())
}
