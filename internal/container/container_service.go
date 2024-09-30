package container

import (
	"context"
	"github.com/secfault-org/hacktober/internal/model/container"
)

type Service interface {
	PullImage(ctx context.Context, image string) error
	StartContainer(ctx context.Context, image string, exposedContainerPort container.Port) (*container.Container, error)
	GetHostPort(ctx context.Context, containerId container.Id, containerPort container.Port) (container.Port, error)
	StopContainer(ctx context.Context, containerId container.Id) error
}
