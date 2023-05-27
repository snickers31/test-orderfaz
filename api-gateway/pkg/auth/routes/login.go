package routes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/snickers31/test-orderfaz/api-gateway/pkg/auth/constants"
	"github.com/snickers31/test-orderfaz/api-gateway/pkg/auth/pb"
)

func Login(ctx *gin.Context, c pb.AuthServiceClient) {
	requestBody := constants.LoginRequestBody{}

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}

	res, err := c.Login(context.Background(), &pb.LoginRequest{
		Msisdn:   requestBody.Msisdn,
		Password: requestBody.Password,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
