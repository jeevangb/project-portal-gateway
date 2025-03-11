package services

import (
	"context"

	"github.com/jeevangb/project-portal-gateway/internal/graph/model"
	"github.com/jeevangb/project-portal-gateway/internal/grpc/proto"
)

func (s *serviceLayer) SignUp(ctx context.Context, input *model.NewUser) (*model.UserResponse, error) {
	req := &proto.SignUpRequest{Name: input.Name, Email: input.Email, Password: input.Password}
	res, err := s.cli.SignUp.SignUp(ctx, req)
	if err != nil {
		return &model.UserResponse{}, err
	}
	resp := convertUserResponse(res)
	return resp, nil
}

func convertUserResponse(res *proto.UserResponse) *model.UserResponse {
	return &model.UserResponse{
		Message: res.Message,
	}
}
