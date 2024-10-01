package service

import (
	"context"
	authv1 "github.com/Pasca11/gRPC-Auth/proto/gen"
	"github.com/Pasca11/justNotes/models"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type gRPCUserService struct {
	conn *grpc.ClientConn
}

func NewGRPCUserService(address string) UserService {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	service := &gRPCUserService{
		conn: conn,
	}
	return service
}

func (s *gRPCUserService) Login(req *models.User) (*models.LoginResponse, error) {
	client := authv1.NewAuthClient(s.conn)

	resp, err := client.Login(context.Background(), &authv1.LoginRequest{})
	if err != nil {
		return nil, err
	}

	return &models.LoginResponse{Token: resp.Token}, nil
}

func (s *gRPCUserService) Register(req *models.User) (*models.User, error) {
	client := authv1.NewAuthClient(s.conn)

	resp, err := client.Register(context.Background(), &authv1.RegisterRequest{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	return &models.User{
		ID: int(resp.UserId), //TODO return ID
	}, nil
}
