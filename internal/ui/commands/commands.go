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

type EnterFlagMsg struct{}
type FlagEnteredMsg struct {
	Flag string
}
type EnteringFlagCanceledMsg struct{}

type ChallengeSolvedMsg *challenge.Challenge
type WrongFlagMsg struct{}

type Quitting struct{}
type TeardownMsg struct{}

func ChallengeStarting(challenge *challenge.Challenge) tea.Cmd {
	return func() tea.Msg {
		return ChallengeStaringMsg(challenge)
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

func EnteringFlagCmd() tea.Msg {
	return EnterFlagMsg{}
}

func EnteringFlagCanceled() tea.Msg {
	return EnteringFlagCanceledMsg{}
}

func FlagEntered(submittedFlag string) tea.Cmd {
	return func() tea.Msg {
		return FlagEnteredMsg{submittedFlag}
	}
}

func QuittingCmd(activeChallenge *challenge.ActiveChallenge) tea.Cmd {
	return func() tea.Msg {
		if activeChallenge != nil && activeChallenge.IsRunning() {
			return TeardownMsg{}
		} else {
			return tea.QuitMsg{}
		}
	}
}

func SubmitFlag(submittedFlag string, activeChallenge *challenge.ActiveChallenge) tea.Cmd {
	return func() tea.Msg {
		if submittedFlag == activeChallenge.Flag {
			return ChallengeSolvedMsg(activeChallenge.Challenge)
		} else {
			return WrongFlagMsg{}
		}
	}
}
