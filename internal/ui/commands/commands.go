package commands

import (
	"context"

	tea "charm.land/bubbletea/v2"
	"github.com/secfault-org/hacktober/internal/model/challenge"
	"github.com/secfault-org/hacktober/internal/session"
)

type SelectChallengeMsg *challenge.Challenge
type GoBackMsg struct{}

type ChallengeStaringMsg *challenge.Challenge
type ChallengeStartedMsg *challenge.ActiveChallenge
type ChallengeStoppingMsg *challenge.ActiveChallenge
type ChallengeStoppedMsg *challenge.Challenge
type ContainerErrorMsg error
type ChallengeTimedOutMsg struct{}

type ActiveChallengeChangedMsg *challenge.ActiveChallenge

type EnterFlagMsg struct{}
type FlagEnteredMsg struct {
	Flag string
}
type EnteringFlagCanceledMsg struct{}

type ChallengeSolvedMsg *challenge.Challenge
type WrongFlagMsg struct{}

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

func SubmitFlag(sess *session.Session, flag string) tea.Cmd {
	return func() tea.Msg {
		if !sess.CheckFlag(flag) {
			return WrongFlagMsg{}
		}
		active := sess.ActiveChallenge()
		if active == nil {
			return nil
		}

		return ChallengeSolvedMsg(active.Challenge)
	}
}

func StartChallenge(sess *session.Session, ch *challenge.Challenge) tea.Cmd {
	return func() tea.Msg {
		if !sess.SetStarting(ch) {
			return nil
		}

		return ChallengeStaringMsg(ch)
	}
}

func SpawnContainer(sess *session.Session, ctx context.Context, ch *challenge.Challenge) tea.Cmd {
	return func() tea.Msg {
		ac, err := sess.Start(ctx, ch)
		if err != nil {
			return ContainerErrorMsg(err)
		}

		return ChallengeStartedMsg(ac)
	}
}

func StopChallenge(sess *session.Session, ctx context.Context) tea.Cmd {
	return func() tea.Msg {
		active := sess.ActiveChallenge()
		if active == nil {
			return nil
		}
		ch := active.Challenge
		if err := sess.Stop(ctx); err != nil {
			return ContainerErrorMsg(err)
		}
		return ChallengeStoppedMsg(ch)
	}
}
