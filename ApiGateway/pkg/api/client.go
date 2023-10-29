package api

import (
	"fmt"

	"github.com/akshayUr04/totalitycorp-apigateway/pkg/config"
	"github.com/akshayUr04/totalitycorp-apigateway/pkg/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type UserServiceClient struct {
	Client pb.UserServiceClient
}

func InitUserSvc(c *config.Config) pb.UserServiceClient {
	cc, err := grpc.Dial(c.UserSvcPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Could not connect:", err)
	}
	return pb.NewUserServiceClient(cc)
}
