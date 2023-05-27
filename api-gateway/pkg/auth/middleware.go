package auth

import (
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/snickers31/test-orderfaz/api-gateway/pkg/auth/pb"
)

type AuthMiddlewareConfig struct {
	svc *ServiceClient
}

func InitAuthMiddleware(svc *ServiceClient) AuthMiddlewareConfig {
	return AuthMiddlewareConfig{svc: svc}
}

func (am *AuthMiddlewareConfig) AuthRequired(ctx *gin.Context) {
	authorization := ctx.Request.Header.Get("authorization")

	if authorization == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"err": "Unauthorized",
		})
		return
	}

	token := strings.Split(authorization, "Bearer ")

	if len(token) != 2 {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"err": "Unauthorized",
		})
		return
	}

	res, err := am.svc.Client.Validate(context.Background(), &pb.ValidateRequest{
		Token: token[1],
	})

	if err != nil || res.GetStatus() != 0 {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"err": err.Error(),
		})
		return
	}

	ctx.Set("claim_id", res.ClaimId)

	ctx.Next()
}
