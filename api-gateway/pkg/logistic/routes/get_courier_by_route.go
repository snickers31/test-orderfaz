package routes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/snickers31/test-orderfaz/api-gateway/pkg/proto"
)

func GetCourierByRoute(ctx *gin.Context, c proto.LogisticServiceClient) {
	dst_name := ctx.Query("destination_name")
	ogn_name := ctx.Query("origin_name")

	res, err := c.GetCourierByRoute(context.Background(), &proto.RouteParams{
		OriginName:      ogn_name,
		DestinationName: dst_name,
	})

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
