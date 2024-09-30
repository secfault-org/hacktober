package commands

import (
	"github.com/secfault-org/hacktober/internal/model/challenge"
	"github.com/secfault-org/hacktober/internal/model/container"
)

type SelectChallengeMsg challenge.Challenge
type ContainerLoadingMsg struct {
	Loading bool
	Message string
}
type GoBackMsg struct{}
type ContainerSpawnedMsg *container.Container
type ContainerStoppedMsg struct{}
type ContainerErrorMsg error
