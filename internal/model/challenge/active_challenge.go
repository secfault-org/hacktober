package challenge

import (
	"github.com/secfault-org/hacktober/internal/model/container"
)

type ActiveChallenge struct {
	Challenge *Challenge
	Container *container.Container
}

func (r *ActiveChallenge) State() container.State {
	return r.Container.State
}

func (r *ActiveChallenge) IsRunning() bool {
	return r.State() == container.Running
}
