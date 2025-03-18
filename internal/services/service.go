package services

import (
	"context"

	"github.com/jeevangb/project-portal-gateway/internal/clients"
	"github.com/jeevangb/project-portal-gateway/internal/graph/model"
)

type serviceLayer struct {
	cli clients.GrpcClient
}

type Service interface {
	SignUp(ctx context.Context, input *model.NewUser) (*model.UserResponse, error)
	Login(ctx context.Context, username string, password string) (*model.UserResponse, error)
	CreateProject(ctx context.Context, input *model.ProjectInput) (*model.Project, error)
	UpdateProject(ctx context.Context, input *model.UpdateProjectInput) (*model.Project, error)
}

func NewService(client clients.GrpcClient) (Service, error) {
	return &serviceLayer{
		cli: client,
	}, nil
}
