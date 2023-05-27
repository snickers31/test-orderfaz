package logistic

import (
	"fmt"

	"github.com/snickers31/test-orderfaz/api-gateway/pkg/config"
	"github.com/snickers31/test-orderfaz/api-gateway/pkg/logistic/pb"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.LogisticServiceClient
}

func InitServiceClient(c *config.Config) pb.LogisticServiceClient {
	cc, err := grpc.Dial(c.LogisticSvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect logistic services:", err)
	}

	return pb.NewLogisticServiceClient(cc)

}
