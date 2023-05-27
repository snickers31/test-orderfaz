package main

import (
	"fmt"
	"log"
	"net"

	"github.com/snickers31/test-orderfaz/auth-svc/pkg/config"
	"github.com/snickers31/test-orderfaz/auth-svc/pkg/db"
	"github.com/snickers31/test-orderfaz/auth-svc/pkg/pb"
	"github.com/snickers31/test-orderfaz/auth-svc/pkg/services"
	"github.com/snickers31/test-orderfaz/auth-svc/pkg/utils"
	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at load config", err)
	}

	h := db.Init(c.DBUrl)

	jwt := utils.JwtWrapper{
		SecretKey:       c.JWTSecretKey,
		Issuer:          "rizal-test-orderfaz",
		ExpirationHours: 24 * 365,
	}

	lis, err := net.Listen("tcp", c.Port)

	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	fmt.Println("Auth Service Running On ", c.Port)

	s := services.Server{
		H:   h,
		Jwt: jwt,
	}

	grpcServer := grpc.NewServer()
	pb.RegisterAuthServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve auth service:", err)
	}
}
