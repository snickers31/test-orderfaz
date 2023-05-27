package logistic

import (
	"github.com/gin-gonic/gin"
	"github.com/snickers31/test-orderfaz/api-gateway/pkg/auth"
	"github.com/snickers31/test-orderfaz/api-gateway/pkg/config"
	"github.com/snickers31/test-orderfaz/api-gateway/pkg/logistic/routes"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/api/courier")
	routes.Use(a.AuthRequired)

	routes.GET("", svc.GetCouriers)
	routes.GET("/find-by-routes", svc.GetCourierByRoute)

}

func (svc *ServiceClient) GetCouriers(ctx *gin.Context) {
	routes.GetCouriers(ctx, svc.Client)
}
func (svc *ServiceClient) GetCourierByRoute(ctx *gin.Context) {
	routes.GetCourierByRoute(ctx, svc.Client)
}
