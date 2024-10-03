package commands

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/secfault-org/hacktober/internal/model/challenge"
)

type SelectChallengeMsg *challenge.Challenge
type GoBackMsg struct{}

type ChallengeStaringMsg *challenge.Challenge
type ChallengeStartedMsg *challenge.ActiveChallenge
type ChallengeStoppingMsg *challenge.ActiveChallenge
type ChallengeStoppedMsg *challenge.Challenge
type ContainerErrorMsg error

type ActiveChallengeChangedMsg *challenge.ActiveChallenge

func ChallengeStarting(challenge *challenge.Challenge) tea.Cmd {
	return func() tea.Msg {
		return ChallengeStaringMsg(challenge)
	}
}

func ChallengeStarted(activeChallenge *challenge.ActiveChallenge) tea.Cmd {
	return func() tea.Msg {
		return ChallengeStartedMsg(activeChallenge)
	}
}

func ChallengeStopping(activeChallenge *challenge.ActiveChallenge) tea.Cmd {
	return func() tea.Msg {
		return ChallengeStoppingMsg(activeChallenge)
	}
}

func UpdateActiveChallenge(activeChallenge *challenge.ActiveChallenge) tea.Cmd {
	return func() tea.Msg {
		return ActiveChallengeChangedMsg(activeChallenge)
	}
}
