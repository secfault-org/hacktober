package challenge

import (
	"github.com/secfault-org/hacktober/internal/model/container"
)

type ActiveChallenge struct {
	Challenge *Challenge
	Container *container.Container
	Flag      string
}

func (r *ActiveChallenge) State() container.State {
	if r.Container != nil {
		return r.Container.State
	}

	return container.Stopped
}

func (r *ActiveChallenge) IsRunning() bool {
	return r.State() == container.Running
}
