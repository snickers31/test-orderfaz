package routes

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/snickers31/test-orderfaz/api-gateway/pkg/logistic/pb"
)

func GetCouriers(ctx *gin.Context, c pb.LogisticServiceClient) {
	page := ctx.Query("page")
	intPage, err := strconv.Atoi(page)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := c.GetCouriers(context.Background(), &pb.Page{
		Page: int64(intPage),
	})

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, &res)

}
