package main

import (
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/snickers31/test-orderfaz/api-gateway/cmd/docs"
	"github.com/snickers31/test-orderfaz/api-gateway/pkg/auth"
	"github.com/snickers31/test-orderfaz/api-gateway/pkg/config"
	"github.com/snickers31/test-orderfaz/api-gateway/pkg/logistic"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title     Rizal Test
func main() {

	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at load config", err)
	}

	r := gin.Default()
	authSvc := *auth.RegisterRoutes(r, &c)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	logistic.RegisterRoutes(r, &c, &authSvc)

	r.Run(c.Port)

}
