package services

import (
	"context"

	"github.com/jeevangb/project-portal-gateway/internal/clients"
	"github.com/jeevangb/project-portal-gateway/internal/graph/model"
)

type serviceLayer struct {
	cli clients.AuthClient
}

type Service interface {
	SignUp(ctx context.Context, input *model.NewUser) (*model.UserResponse, error)
	Login(ctx context.Context, username string, password string) (*model.UserResponse, error)
}

func NewService(client clients.AuthClient) (Service, error) {
	return &serviceLayer{
		cli: client,
	}, nil
}
