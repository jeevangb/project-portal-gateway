package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.66

import (
	"context"
	"time"

	"github.com/jeevangb/project-portal-gateway/internal/graph/model"
)

// SignUp is the resolver for the signUp field.
func (r *mutationResolver) SignUp(ctx context.Context, input model.NewUser) (*model.UserResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 50*time.Minute)
	defer cancel()
	return r.Service.SignUp(ctx, &input)
}

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, email string, password string) (*model.UserResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 50*time.Minute)
	defer cancel()
	return r.Service.Login(ctx, email, password)
}

// CreateProject is the resolver for the createProject field.
func (r *mutationResolver) CreateProject(ctx context.Context, input model.ProjectInput) (*model.Project, error) {
	ctx, cancel := context.WithTimeout(ctx, 50*time.Minute)
	defer cancel()
	return r.Service.CreateProject(ctx, &input)
}

// UpdateProject is the resolver for the updateProject field.
func (r *mutationResolver) UpdateProject(ctx context.Context, input *model.UpdateProjectInput) (*model.Project, error) {
	ctx, cancel := context.WithTimeout(ctx, 50*time.Minute)
	defer cancel()
	return r.Service.UpdateProject(ctx, input)
}

// DeleteProject is the resolver for the deleteProject field.
func (r *mutationResolver) DeleteProject(ctx context.Context, name string) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, 50*time.Minute)
	defer cancel()
	return r.Service.DeleteProject(ctx, name)
}

// HealthCheck is the resolver for the healthCheck field.
func (r *queryResolver) HealthCheck(ctx context.Context) (*model.HealthStatus, error) {
	return &model.HealthStatus{
		Status:  200,
		Message: "Success",
	}, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
