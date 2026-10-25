package application

import (
	"charm.land/log/v2"
	"github.com/secfault-org/hacktober/internal/config"
	"github.com/secfault-org/hacktober/internal/container"
	"github.com/secfault-org/hacktober/internal/repository"
)

type App struct {
	repo      repository.Repository
	container container.Service
	config    *config.Config
	logger    *log.Logger
}

func NewApp(cfg *config.Config, repo repository.Repository, container container.Service, logger *log.Logger) *App {
	return &App{
		repo:      repo,
		container: container,
		config:    cfg,
		logger:    logger,
	}
}

func (a *App) Repo() repository.Repository         { return a.repo }
func (a *App) ContainerService() container.Service { return a.container }
func (a *App) Config() *config.Config              { return a.config }
func (a *App) Logger() *log.Logger                 { return a.logger }
