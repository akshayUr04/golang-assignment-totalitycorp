package routes

import (
	"context"
	"net/http"
	"strconv"

	"github.com/akshayUr04/totalitycorp-apigateway/pkg/pb"
	"github.com/akshayUr04/totalitycorp-apigateway/pkg/responce"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

func GetUserById(ctx *gin.Context, p pb.UserServiceClient) {
	paramsId := ctx.Param("id")
	id, err := strconv.Atoi(paramsId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "can't get userID",
		})
		return
	}
	if id <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "id canot be zero or less",
		})
		return
	}
	result, err := p.GetUserById(context.Background(), &pb.GetUserRequest{
		Id: int32(id),
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "cant get responce form the server",
		})
		return
	}
	resp := []responce.UserResponce{}
	copier.Copy(&resp, &result)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "details found successfully",
		"data":    resp,
	})
}

func GetUserByIds(ctx *gin.Context, p pb.UserServiceClient) {
	var ids []int32
	queryIds := ctx.QueryArray("id")
	for _, val := range queryIds {
		id, err := strconv.Atoi(val)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error":   err.Error(),
				"message": "can't get userID",
			})
			return
		}
		if id <= 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "id canot be zero or less",
			})
			return
		}
		ids = append(ids, int32(id))
	}
	result, err := p.GetUsersByIds(context.Background(), &pb.GetUsersRequest{
		Ids: ids,
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "cant get responce form the server",
		})
		return
	}
	resp := make([]responce.UserResponce, len(result.Users))
	copier.Copy(&resp, &result.Users)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "details found successfully",
		"data":    resp,
	})

}
