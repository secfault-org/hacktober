package podman

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/secfault-org/hacktober/internal/model/container"
	nettypes "go.podman.io/common/libnetwork/types"
	"go.podman.io/podman/v6/pkg/bindings"
	"go.podman.io/podman/v6/pkg/bindings/containers"
	"go.podman.io/podman/v6/pkg/bindings/images"
	"go.podman.io/podman/v6/pkg/specgen"
)

type Service struct{}

func NewContainerService() *Service { return &Service{} }

func (s *Service) Connect(ctx context.Context) (context.Context, error) {
	sock_dir := os.Getenv("XDG_RUNTIME_DIR")
	if sock_dir == "" {
		sock_dir = "/var/run"
	}
	socket := "unix:" + sock_dir + "/podman/podman.sock"

	// Connect to Podman socket
	ctx, err := bindings.NewConnection(ctx, socket)
	if err != nil {
		return nil, err
	}
	return ctx, nil
}

func (s *Service) PullImage(ctx context.Context, image string) error {
	policy := "missing"
	_, err := images.Pull(ctx, image, &images.PullOptions{Policy: &policy})
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) StartContainer(ctx context.Context, image, flag string, exposedContainerPort container.Port) (*container.Container, error) {
	if err := s.PullImage(ctx, image); err != nil {
		return nil, fmt.Errorf("pull %s: %w", image, err)
	}
	spec := specgen.NewSpecGenerator(image, false)
	spec.PortMappings = []nettypes.PortMapping{
		{
			HostPort:      0, // Let Podman assign a random available port
			ContainerPort: exposedContainerPort,
			Protocol:      "tcp",
		},
	}
	spec.Env = map[string]string{
		"FLAG": flag,
	}
	spec.SeccompProfilePath = "unconfined"
	remove := true
	spec.Remove = &remove
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

	port, err := s.GetHostPort(ctx, r.ID, exposedContainerPort)
	if err != nil {
		return nil, err
	}

	return &container.Container{
		ID:       r.ID,
		HostPort: port,
		State:    container.Running,
	}, nil
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
	}

	return 0, fmt.Errorf("port %d not found for container %s", containerPort, containerId)
}

func (s *Service) StopContainer(ctx context.Context, containerId container.Id) error {
	timeout := uint(3)
	return containers.Stop(ctx, containerId, &containers.StopOptions{Timeout: &timeout})
}
