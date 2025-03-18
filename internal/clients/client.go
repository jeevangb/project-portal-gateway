package clients

import (
	"errors"
	"fmt"
	"time"

	"github.com/jeevangb/project-portal-gateway/internal/config"
	"github.com/jeevangb/project-portal-gateway/internal/grpc/proto"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

type GrpcClient struct {
	UserService   proto.AuthServiceClient
	MentorService proto.ProjectServiceClient
}

func GetGrpcAuthServernection(cfg config.Config) (GrpcClient, error) {
	log.Info().Msg("gRPC connection initiated for auth service")

	ServerPort := cfg.AuthPORT
	if ServerPort == "" {
		err := errors.New("gRPC server port is empty")
		log.Error().Err(err).Msg("gRPC server port is empty")
		return GrpcClient{}, err
	}

	// Create a connection to the gRPC server
	conn, err := grpc.Dial(ServerPort, grpc.WithTimeout(time.Minute*5), grpc.WithInsecure())
	if err != nil {
		log.Error().Err(fmt.Errorf("unable to connect to %v: %w", ServerPort, err))
		return GrpcClient{}, err // Return the error instead of creating an empty client
	}

	// Ensure the connection is closed when done
	log.Info().Msg("Connection established with auth service at port: " + ServerPort)
	return GrpcClient{
		UserService:   proto.NewAuthServiceClient(conn),
		MentorService: proto.NewProjectServiceClient(conn),
	}, nil
}
