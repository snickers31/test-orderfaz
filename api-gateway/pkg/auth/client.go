package auth

import (
	"fmt"

	"github.com/snickers31/test-orderfaz/api-gateway/pkg/config"
	"github.com/snickers31/test-orderfaz/api-gateway/pkg/proto"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client proto.AuthServiceClient
}

func InitServiceClient(c *config.Config) proto.AuthServiceClient {
	cc, err := grpc.Dial(c.AuthSvcUrl, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Could not connect to auth service: ", err)
	}

	return proto.NewAuthServiceClient(cc)

}
