package services

import (
	"context"

	"github.com/snickers31/test-orderfaz/logistic-svc/pkg/db"
	"github.com/snickers31/test-orderfaz/logistic-svc/pkg/helpers"
	"github.com/snickers31/test-orderfaz/logistic-svc/pkg/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	pb.UnimplementedLogisticServiceServer
	H db.Handler
}

func (s *Server) GetCouriers(ctx context.Context, req *pb.Page) (*pb.Couriers, error) {
	var page int64 = 1

	if req.GetPage() != 0 {
		page = req.GetPage()
	}

	var pagination pb.Pagination
	var couriers []*pb.Courier

	sql := s.H.DB.Table("couriers").Select("logistic_name", "amount", "destination_name", "origin_name", "duration")
	offset, limit := helpers.Pagination(sql, page, &pagination)

	rows, err := sql.Offset(int(offset)).Limit(int(limit)).Rows()

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	for rows.Next() {
		var courier pb.Courier

		if err := rows.Scan(&courier.LogisticName, &courier.Amount, &courier.DestinationName, &courier.OriginName, &courier.Duration); err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		couriers = append(couriers, &courier)
	}

	response := &pb.Couriers{
		Pagination: &pagination,
		Data:       couriers,
	}

	return response, nil

}
func (s *Server) GetCourierByRoute(ctx context.Context, req *pb.RouteParams) (*pb.Courier, error) {
	row := s.H.DB.Table("couriers AS c").Select("logistic_name", "amount", "destination_name", "origin_name", "duration").
		Where("c.origin_name = ? AND c.destination_name = ? ", req.GetOriginName(), req.GetDestinationName()).Row()

	var courier pb.Courier

	if err := row.Scan(&courier.LogisticName, &courier.Amount, &courier.DestinationName, &courier.OriginName, &courier.Duration); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &courier, nil

}
