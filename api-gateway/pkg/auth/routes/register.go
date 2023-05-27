package routes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/snickers31/test-orderfaz/api-gateway/pkg/auth/constants"
	"github.com/snickers31/test-orderfaz/api-gateway/pkg/auth/pb"
)

func Register(ctx *gin.Context, c pb.AuthServiceClient) {
	requestBody := constants.RegisterRequestBody{}

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}

	res, _ := c.Register(context.Background(), &pb.RegisterRequest{
		Msisdn:   requestBody.Msisdn,
		Name:     requestBody.Name,
		Username: requestBody.Username,
		Password: requestBody.Password,
	})

	if res.GetStatus() != 0 {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"err": res.GetError(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Akun berhasil dibuat.",
	})

}
