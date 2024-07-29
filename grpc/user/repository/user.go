package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	pb "user/proto"
)

func (r *Repository) GetUserList(ctx context.Context) (*[]pb.GetUserResponse, error) {
	var filter bson.M
	var users *[]pb.GetUserResponse
	cursor, err := r.userRef.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	if err = cursor.All(context.TODO(), &users); err != nil {
		return nil, err
	}

	return users, nil
}
