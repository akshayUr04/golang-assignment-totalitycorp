package service

import (
	"context"
	"fmt"

	mockdb "github.com/akshayUr04/totalitycorp-service/pkg/mockDb"
	"github.com/akshayUr04/totalitycorp-service/pkg/pb"
)

var db = mockdb.Users

type UserServer struct {
	pb.UnimplementedUserServiceServer
}

func (s *UserServer) GetUserById(ctx context.Context, req *pb.GetUserRequest) (*pb.User, error) {
	users, ok := db[req.Id]
	if !ok {
		return nil, fmt.Errorf("no user coresponding to the given ID")
	}
	return &pb.User{
		Id:      users.ID,
		Fname:   users.Fname,
		City:    users.City,
		Phone:   users.Phone,
		Height:  users.Height,
		Married: users.Married,
	}, nil
}

func (s *UserServer) GetUsersByIds(ctx context.Context, req *pb.GetUsersRequest) (*pb.UserList, error) {
	response := &pb.UserList{
		Users: make([]*pb.User, 0, len(req.Ids)),
	}

	for _, userID := range req.Ids {
		user, ok := mockdb.Users[userID]
		if !ok {
			return nil, fmt.Errorf("user with ID %d not found", userID)
		}

		pbUser := &pb.User{
			Id:      user.ID,
			Fname:   user.Fname,
			City:    user.City,
			Phone:   user.Phone,
			Height:  user.Height,
			Married: user.Married,
		}

		response.Users = append(response.Users, pbUser)
	}

	return response, nil

}
