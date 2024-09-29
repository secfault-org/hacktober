package backend

import (
	"context"
	"github.com/charmbracelet/log"
	"github.com/secfault-org/hacktober/pkg/container"
	"github.com/secfault-org/hacktober/pkg/repository"
)

type Backend struct {
	Repo             repository.Repository
	ContainerService container.Service
	logger           *log.Logger
}

func NewBackend(ctx context.Context, repo repository.Repository, containerService container.Service) *Backend {
	logger := log.FromContext(ctx).WithPrefix("backend")
	return &Backend{
		Repo:             repo,
		ContainerService: containerService,
		logger:           logger,
	}
}
