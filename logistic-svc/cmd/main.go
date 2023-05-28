package main

import (
	"fmt"
	"log"
	"net"

	"github.com/snickers31/test-orderfaz/logistic-svc/pkg/config"
	"github.com/snickers31/test-orderfaz/logistic-svc/pkg/db"
	"github.com/snickers31/test-orderfaz/logistic-svc/pkg/pb"
	"github.com/snickers31/test-orderfaz/logistic-svc/pkg/services"
	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at load config", err)
	}

	h := db.Init(c.DBUrl)
	lis, err := net.Listen("tcp", c.Port)

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	fmt.Println("logistic Service Running On ", c.Port)

	s := services.Server{
		H: h,
	}

	grpcServer := grpc.NewServer()
	pb.RegisterLogisticServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
