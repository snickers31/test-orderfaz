package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/snickers31/test-orderfaz/api-gateway/pkg/config"
)

func main() {

	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at load config", err)
	}

	r := gin.Default()

	r.Run(c.Port)

}
