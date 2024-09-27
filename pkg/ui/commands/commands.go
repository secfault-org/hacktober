package commands

import "github.com/secfault-org/hacktober/pkg/model"

type SelectChallengeMsg model.Challenge
type ContainerLoadingMsg struct {
	Loading bool
	Message string
}
type GoBackMsg struct{}
type ContainerSpawnedMsg struct {
	model.Challenge
	State int
}
