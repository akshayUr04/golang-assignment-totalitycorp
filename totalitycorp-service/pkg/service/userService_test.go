package service

import (
	"context"
	"testing"

	"github.com/akshayUr04/totalitycorp-service/pkg/pb"
	"github.com/stretchr/testify/assert"
)

func TestGetUserById(t *testing.T) {
	server := &UserServer{}

	// Test success case.
	req := &pb.GetUserRequest{Id: 1}
	resp, err := server.GetUserById(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, &pb.User{
		Id:      1,
		Fname:   "Akshay",
		City:    "Kozhikode",
		Phone:   1234567890,
		Height:  5.8,
		Married: false,
	}, resp)

	// Test failure case.
	req = &pb.GetUserRequest{Id: 100}
	resp, err = server.GetUserById(context.Background(), req)
	assert.Error(t, err)
	assert.Equal(t, "user with ID 100 not found", err.Error())
	assert.Nil(t, resp)
}

func TestGetUsersByIds(t *testing.T) {
	server := &UserServer{}

	// Test success case.
	req := &pb.GetUsersRequest{Ids: []int32{1, 2, 3}}
	resp, err := server.GetUsersByIds(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, &pb.UserList{
		Users: []*pb.User{
			{
				Id:      1,
				Fname:   "Akshay",
				City:    "Kozhikode",
				Phone:   1234567890,
				Height:  5.8,
				Married: false,
			},
			{
				Id:      2,
				Fname:   "Amal",
				City:    "Pathanamthitta",
				Phone:   1234567890,
				Height:  5.4,
				Married: true,
			},
			{
				Id:      3,
				Fname:   "Amar",
				City:    "Palakkad",
				Phone:   1234567890,
				Height:  5.3,
				Married: true,
			},
		},
	}, resp)

	// Test failure case.
	req = &pb.GetUsersRequest{Ids: []int32{100, 200, 300}}
	resp, err = server.GetUsersByIds(context.Background(), req)
	assert.Error(t, err)
	assert.Equal(t, "user with ID 100 not found", err.Error())
	assert.Nil(t, resp)
}
