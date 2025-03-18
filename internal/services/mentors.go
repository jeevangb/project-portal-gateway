package services

import (
	"context"

	"github.com/jeevangb/project-portal-gateway/internal/graph/model"
	"github.com/jeevangb/project-portal-gateway/internal/grpc/proto"
)

func (s *serviceLayer) CreateProject(ctx context.Context, input *model.ProjectInput) (*model.Project, error) {
	// techStackSlice := strings.Split(input.TechnologyStack, ",")
	req := &proto.CreateProjectRequest{
		Name:            input.Name,
		Description:     input.Description,
		TechnologyStack: input.TechnologyStack,
		MentorName:      input.MentorName,
		Status:          input.Status}
	res, err := s.cli.MentorService.CreateProject(ctx, req)
	if err != nil {
		return nil, err
	}
	resp := convertResponse(res)
	return &resp, nil
}

func (s *serviceLayer) UpdateProject(ctx context.Context, input *model.UpdateProjectInput) (*model.Project, error) {
	req := &proto.UpdateProjectRequest{
		Name:            *input.Title,
		Description:     *input.Description,
		TechnologyStack: input.TechnologyStack,
		MentorName:      *input.Description,
		Status:          *input.Status,
	}
	res, err := s.cli.MentorService.UpdateProject(ctx, req)
	if err != nil {
		return nil, err
	}
	resp := convertResponse(res)
	return &resp, nil
}

func convertResponse(res *proto.Project) model.Project {
	// techStack := strings.Split(res.GetTechnologyStack, ",")
	return model.Project{
		Name:            res.GetName(),
		Description:     res.GetDescription(),
		TechnologyStack: res.GetTechnologyStack(),
		Status:          res.GetStatus(),
	}
}
