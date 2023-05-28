package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/snickers31/test-orderfaz/api-gateway/pkg/auth/routes"
	"github.com/snickers31/test-orderfaz/api-gateway/pkg/config"
)

func RegisterRoutes(r *gin.Engine, c *config.Config) *ServiceClient {

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	a := InitAuthMiddleware(svc)

	route := r.Group("/api/auth")
	route.POST("/register", svc.Register)
	route.POST("/login", svc.Login)

	route.Use(a.AuthRequired)

	route.GET("/get-claim-id", svc.GetClaimId)

	return svc
}

func (svc *ServiceClient) Register(ctx *gin.Context) {
	routes.Register(ctx, svc.Client)
}

func (svc *ServiceClient) Login(ctx *gin.Context) {
	routes.Login(ctx, svc.Client)
}
func (svc *ServiceClient) GetClaimId(ctx *gin.Context) {
	routes.GetClaimsId(ctx)
}
