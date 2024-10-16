package podman

import (
	"context"
	"fmt"
	nettypes "github.com/containers/common/libnetwork/types"
	"github.com/containers/podman/v5/pkg/bindings"
	"github.com/containers/podman/v5/pkg/bindings/containers"
	"github.com/containers/podman/v5/pkg/bindings/images"
	"github.com/containers/podman/v5/pkg/specgen"
	"github.com/secfault-org/hacktober/internal/model/container"
	"os"
	"runtime"
	"strconv"
)

type Service struct {
	runningContainer *container.Container
}

func NewContainerService(ctx context.Context) *Service {
	service := &Service{
		runningContainer: nil,
	}

	// TODO: Find a better way to handle this
	runtime.SetFinalizer(service, func(s *Service) {
		if s.runningContainer != nil {
			ctx, _ := s.Connect(context.Background())
			_ = s.StopContainer(ctx, s.runningContainer.ID)
		}
	})

	return service
}

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
	_, err := images.Pull(ctx, image, &images.PullOptions{})
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) StartContainer(ctx context.Context, image, flag string, exposedContainerPort container.Port) (*container.Container, error) {
	if s.runningContainer != nil {
		return nil, fmt.Errorf("container already running")
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

	s.runningContainer = &container.Container{
		ID:       r.ID,
		HostPort: port,
		State:    container.Running,
	}

	return s.runningContainer, nil
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
	if s.runningContainer == nil {
		return fmt.Errorf("container not running")
	}
	s.runningContainer.State = container.Stopping
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
	s.runningContainer = nil
	return nil
}
