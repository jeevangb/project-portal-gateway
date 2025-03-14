package services

import (
	"context"

	"github.com/jeevangb/project-portal-gateway/internal/auth"
	"github.com/jeevangb/project-portal-gateway/internal/graph/model"
	"github.com/jeevangb/project-portal-gateway/internal/grpc/proto"
)

func (s *serviceLayer) SignUp(ctx context.Context, input *model.NewUser) (*model.UserResponse, error) {
	req := &proto.SignUpRequest{Name: input.Name, Email: input.Email, Password: input.Password}
	res, err := s.cli.UserService.SignUp(ctx, req)
	if err != nil {
		return &model.UserResponse{}, err
	}
	resp := convertUserResponse(res, "")
	return resp, nil
}

func (s *serviceLayer) Login(ctx context.Context, email string, password string) (*model.UserResponse, error) {
	req := &proto.LoginRequest{Email: email, Password: password}
	res, err := s.cli.UserService.Login(ctx, req)
	if err != nil {
		return &model.UserResponse{}, nil
	}
	token, err := auth.GenerateJWT(email)
	if err != nil {
		return nil, err
	}
	resp := convertUserResponse(res, token)
	return resp, nil
}

func convertUserResponse(res *proto.UserResponse, token string) *model.UserResponse {
	return &model.UserResponse{
		Message: res.Message,
		Token:   &token,
	}
}
