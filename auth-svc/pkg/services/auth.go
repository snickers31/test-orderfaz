package services

import (
	"context"
	"fmt"
	"regexp"

	"github.com/google/uuid"
	"github.com/snickers31/test-orderfaz/auth-svc/pkg/db"
	"github.com/snickers31/test-orderfaz/auth-svc/pkg/models"
	"github.com/snickers31/test-orderfaz/auth-svc/pkg/pb"
	"github.com/snickers31/test-orderfaz/auth-svc/pkg/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	pb.UnimplementedAuthServiceServer
	H   db.Handler
	Jwt utils.JwtWrapper
}

func (s *Server) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {

	var user models.User

	pattern := "^62"

	match, _ := regexp.MatchString(pattern, req.GetMsisdn())

	if !match {
		return &pb.RegisterResponse{
			Status: int64(codes.Aborted),
			Error:  "kolom MSISDN harus diawali oleh 62",
		}, fmt.Errorf("kolom MSISDN harus diawali oleh 62")
	}

	var count int64

	err := s.H.DB.Model(&models.User{}).Where("msisdn = ? OR username = ?", req.GetMsisdn(), req.GetUsername()).Count(&count).Error

	if err != nil {
		return &pb.RegisterResponse{
			Status: int64(codes.Internal),
			Error:  err.Error(),
		}, err
	}

	if count > 0 {
		return &pb.RegisterResponse{
			Status: int64(codes.AlreadyExists),
			Error:  "data MSISDN atau username telah terdaftar",
		}, fmt.Errorf("data MSISDN atau username telah terdaftar")
	}

	user.ID = uuid.New()
	user.MSISDN = req.GetMsisdn()
	user.Name = req.GetName()
	user.Username = req.GetUsername()
	user.Password = utils.HashPassword(req.GetPassword())

	if err := s.H.DB.Create(&user).Error; err != nil {
		return &pb.RegisterResponse{
			Status: int64(codes.Internal),
			Error:  err.Error(),
		}, err
	}

	return &pb.RegisterResponse{
		Status: int64(codes.OK),
		Error:  "",
	}, nil
}
func (s *Server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {

	var user models.User
	if result := s.H.DB.Table("users").Where("msisdn = ? ", req.GetMsisdn()).First(&user); result.Error != nil {
		return nil, status.Error(codes.AlreadyExists, "User not found")
	}

	match := utils.CheckPasswordHash(req.GetPassword(), user.Password)

	if !match {
		return nil, status.Error(codes.NotFound, "Password not match")
	}

	token, _ := s.Jwt.GenerateToken(user)

	return &pb.LoginResponse{
		Status: int64(codes.OK),
		Token:  token,
	}, nil

}

func (s *Server) Validate(ctx context.Context, req *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	claims, err := s.Jwt.ValidateToken(req.GetToken())

	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	var user models.User
	if result := s.H.DB.Table("users").Where("id = ? ", claims.Id).First(&user); result.Error != nil {
		return nil, status.Error(codes.NotFound, "User not found")
	}

	return &pb.ValidateResponse{
		Status:  int64(codes.OK),
		ClaimId: user.ID.String(),
	}, nil

}
