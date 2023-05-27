package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetClaimsId(ctx *gin.Context) {

	claimId, _ := ctx.Get("claim_id")

	ctx.JSON(http.StatusOK, gin.H{
		"claim_id": claimId,
	})

}
