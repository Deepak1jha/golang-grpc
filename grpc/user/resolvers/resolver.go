package resolvers

import (
	pb "user/proto"
	"user/repository"
)

type Resolver struct {
	pb.UnimplementedUserServiceServer
	repository *repository.Repository
}

// func (r Resolver) GetUser(ctx context.Context, request *pb.GetUserRequest) (*pb.GetUserResponse, error) {
// 	//TODO implement me
// 	panic("implement me")
// }

func (r Resolver) mustEmbedUnimplementedUserServiceServer() {
	//TODO implement me
	panic("implement me")
}

func NewResolver(repository *repository.Repository) *Resolver {
	return &Resolver{
		repository: repository,
	}
}
