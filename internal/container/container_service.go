package container

import "context"

type Id = string
type Port = uint16

type Service interface {
	PullImage(ctx context.Context, image string) error
	StartContainer(ctx context.Context, image string, exposedContainerPort Port) (*Id, error)
	GetHostPort(ctx context.Context, containerId Id, containerPort Port) (Port, error)
	StopContainer(ctx context.Context, containerId Id) error
}
