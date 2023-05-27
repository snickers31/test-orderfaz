package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/snickers31/test-orderfaz/api-gateway/pkg/auth"
	"github.com/snickers31/test-orderfaz/api-gateway/pkg/config"
	"github.com/snickers31/test-orderfaz/api-gateway/pkg/logistic"
)

func main() {

	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at load config", err)
	}

	r := gin.Default()
	authSvc := *auth.RegisterRoutes(r, &c)

	logistic.RegisterRoutes(r, &c, &authSvc)

	r.Run(c.Port)

}
