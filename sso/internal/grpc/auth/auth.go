package auth

import (
	"context"
	g1 "sso/protos/gen/go/sso"
	"sso/sso/internal/validator"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Auth interface {
	Login(ctx context.Context, email string, password string, appID int) (token string, err error)
	RegisterNewUser(ctx context.Context, email, password string) (userID int64, err error)
	IsAdmin(ctx context.Context, userID int64) (bool, error)
}

type serverAPI struct {
	g1.UnimplementedAuthServer
	auth Auth
}

func Register(gRPC *grpc.Server, auth Auth) {
	g1.RegisterAuthServer(gRPC, &serverAPI{auth: auth})
}

// TODO: ADD serverErrorResponses, with error.go file helper
func (s *serverAPI) Login(ctx context.Context, req *g1.LoginRequest) (*g1.LoginResponse, error) {

	v := validator.New()

	if err := validateLogin(req, v); err != nil {
		return nil, err
	}
	token, err := s.auth.Login(ctx, req.Email, req.Password, int(req.GetAppId()))
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &g1.LoginResponse{
		Token: token,
	}, nil
	//TODO :Implement auth service

}
func (s *serverAPI) Register(ctx context.Context, req *g1.RegisterRequest) (*g1.RegisterResponse, error) {
	v := validator.New()
	if err := validateRegister(req, v); err != nil {
		return nil, err
	}
	userID, err := s.auth.RegisterNewUser(ctx, req.GetEmail(), req.GetPassword())
	if err != nil {
		//TODO :: ...
		return nil, err
	}
	return &g1.RegisterResponse{
		UserId: userID,
	}, nil

}

func (s *serverAPI) IsAdmin(ctx context.Context, req *g1.IsAdminRequest) (*g1.IsAdminResponse, error) {
	panic("implement me")
}

func validateLogin(req *g1.LoginRequest, v *validator.Validator) error {
	v.Check(validator.Matches(req.Email, validator.EmailRX), "Email", "Should be correct email address")
	v.Check(req.AppId <= 0, "App ID", "Should be greater than zero")
	v.Check(req.Password != "", "password", "must be provided")
	v.Check(len(req.Password) >= 2, "password", "must be at least 2 bytes long")
	v.Check(len(req.Password) <= 72, "password", "must be no more than 72 bytes long")
	if !v.Valid() {
		return status.Errorf(codes.Internal, "validateLogin:problem")
	}
	return nil
}

func validateRegister(req *g1.RegisterRequest, v *validator.Validator) error {
	v.Check(validator.Matches(req.Email, validator.EmailRX), "Email", "Should be correct email address")
	v.Check(req.Password != "", "password", "must be provided")
	v.Check(len(req.Password) >= 2, "password", "must be at least 2 bytes long")
	v.Check(len(req.Password) <= 72, "password", "must be no more than 72 bytes long")
	if !v.Valid() {
		return status.Errorf(codes.Internal, "validateLogin:problem")
	}
	return nil
}
