package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"user/connection"
	pb "user/proto" // Import the generated protobuf code
	"user/repository"
	"user/resolvers"
)

// server implements the UserServiceServer interface
type server struct {
	pb.UnimplementedUserServiceServer
}

var (
	resolver    *resolvers.Resolver
	connections *connection.Connection
	err         error
	done        chan bool
)

func init() {
	mongoConn := connection.MongoConnect("mongodb://localhost:27017/")
	connections := connection.Connection{Mongo: mongoConn}
	newRepository := repository.NewRepository(&connections)
	resolver = resolvers.NewResolver(newRepository)
}
func main() {

	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		return
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, resolver)
	log.Println("Server is running on port 50051")
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
