package api

import (
	"fmt"

	"github.com/akshayUr04/totalitycorp-apigateway/pkg/api/routes"
	"github.com/akshayUr04/totalitycorp-apigateway/pkg/config"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, cfg *config.Config) {
	svc := &UserServiceClient{
		Client: InitUserSvc(cfg),
	}

	r.GET("/user/:id", svc.GetUserById)
	r.GET("/users/", svc.GetUserByIds)

}

func (s *UserServiceClient) GetUserById(ctx *gin.Context) {
	routes.GetUserById(ctx, s.Client)
}

func (s *UserServiceClient) GetUserByIds(ctx *gin.Context) {
	fmt.Println("API Gateway : GetUserByIds")
	routes.GetUserByIds(ctx, s.Client)
}
