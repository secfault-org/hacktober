package backend

import (
	"github.com/charmbracelet/log"
	"github.com/secfault-org/hacktober/internal/config"
	"github.com/secfault-org/hacktober/internal/container"
	"github.com/secfault-org/hacktober/internal/repository"
)

type Backend struct {
	Config           *config.Config
	Repo             repository.Repository
	ContainerService container.Service
	logger           *log.Logger
}

func NewBackend(
	config *config.Config,
	repo repository.Repository,
	logger *log.Logger,
	containerService container.Service,
) *Backend {
	return &Backend{
		Repo:             repo,
		ContainerService: containerService,
		logger:           logger,
		Config:           config,
	}
}

func (b *Backend) Logger() *log.Logger {
	return b.logger
}
