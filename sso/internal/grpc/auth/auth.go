package auth

import (
	"context"
	g1 "sso/protos/gen/go/sso"

	"google.golang.org/grpc"
)

type serverAPI struct {
	g1.UnimplementedAuthServer
}

func Register(gRPC *grpc.Server) {
	g1.RegisterAuthServer(gRPC, &serverAPI{})
}

func (s *serverAPI) Login(ctx context.Context, req *g1.LoginRequest) (*g1.LoginResponse, error) {
	return &g1.LoginResponse{Token: req.GetEmail()}, nil
}
func (s *serverAPI) Register(ctx context.Context, req *g1.RegisterRequest) (*g1.RegisterResponse, error) {
	panic("implement me")
}

func (s *serverAPI) IsAdmin(ctx context.Context, req *g1.IsAdminRequest) (*g1.IsAdminResponse, error) {
	panic("implement me")
}
