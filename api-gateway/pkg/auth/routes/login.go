package routes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/snickers31/test-orderfaz/api-gateway/pkg/auth/constants"
	"github.com/snickers31/test-orderfaz/api-gateway/pkg/proto"
)

// PostBook             godoc
// @Summary      Store a new book
// @Description  Takes a book JSON and store in DB. Return saved JSON.
// @Tags         books
// @Produce      json
func Login(ctx *gin.Context, c proto.AuthServiceClient) {
	requestBody := constants.LoginRequestBody{}

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}

	res, err := c.Login(context.Background(), &proto.LoginRequest{
		Msisdn:   requestBody.Msisdn,
		Password: requestBody.Password,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
