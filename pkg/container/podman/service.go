package podman

import (
	"context"
	"fmt"
	nettypes "github.com/containers/common/libnetwork/types"
	"github.com/containers/podman/v5/pkg/bindings/containers"
	"github.com/containers/podman/v5/pkg/bindings/images"
	"github.com/containers/podman/v5/pkg/specgen"
	"github.com/secfault-org/hacktober/pkg/container"
	"strconv"
)

type Service struct {
}

func NewContainerService(ctx context.Context) *Service {
	return &Service{}
}

func (s *Service) PullImage(ctx context.Context, image string) error {
	_, err := images.Pull(ctx, image, &images.PullOptions{})
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) StartContainer(ctx context.Context, image string, exposedContainerPort container.Port) (*container.Id, error) {
	spec := specgen.NewSpecGenerator(image, false)
	spec.PortMappings = []nettypes.PortMapping{
		{
			HostPort:      0, // Let Podman assign a random available port
			ContainerPort: exposedContainerPort,
			Protocol:      "tcp",
		},
	}
	r, err := containers.CreateWithSpec(ctx, spec, &containers.CreateOptions{})
	if err != nil {
		return nil, err
	}

	err = containers.Start(ctx, r.ID, &containers.StartOptions{})
	if err != nil {
		return nil, err
	}

	_, err = containers.Wait(ctx, r.ID, &containers.WaitOptions{
		Conditions: []string{"running"},
	})
	if err != nil {
		return nil, err
	}
	return &r.ID, nil
}

func (s *Service) GetHostPort(ctx context.Context, containerId container.Id, containerPort container.Port) (container.Port, error) {
	ctrData, err := containers.Inspect(ctx, containerId, &containers.InspectOptions{})
	if err != nil {
		return 0, err
	}
	portInfo := ctrData.NetworkSettings.Ports[fmt.Sprintf("%d/tcp", containerPort)]
	if len(portInfo) > 0 {
		assignedPort, err := strconv.ParseUint(portInfo[0].HostPort, 10, 16)
		if err != nil {
			return 0, err
		}
		return uint16(assignedPort), nil
	} else {
		return 0, fmt.Errorf("port %d not found for container %s", containerPort, containerId)
	}
}

func (s *Service) StopContainer(ctx context.Context, containerId container.Id) error {
	err := containers.Stop(ctx, containerId, &containers.StopOptions{})
	if err != nil {
		return err
	}
	_, err = containers.Wait(ctx, containerId, &containers.WaitOptions{
		Conditions: []string{"stopped"},
	})
	if err != nil {
		return err
	}
	return nil
}
