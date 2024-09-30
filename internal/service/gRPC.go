package service

import (
	"context"
	authv1 "github.com/Pasca11/gRPC-Auth/proto/gen"
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

func (s *gRPCUserService) Login(req *authv1.LoginRequest) (*authv1.LoginResponse, error) {
	client := authv1.NewAuthClient(s.conn)

	resp, err := client.Login(context.Background(), &authv1.LoginRequest{})
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *gRPCUserService) Register(req *authv1.RegisterRequest) (*authv1.RegisterResponse, error) {
	client := authv1.NewAuthClient(s.conn)

	resp, err := client.Register(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
