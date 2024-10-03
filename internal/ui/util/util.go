package util

import (
	"fmt"
	"github.com/secfault-org/hacktober/internal/model/challenge"
)

func ActiveChallengeStatusMessage(activeChallenge *challenge.ActiveChallenge) string {
	if activeChallenge == nil || activeChallenge.Container == nil {
		return ""
	}

	title := activeChallenge.Challenge.Name
	message := activeChallenge.State()
	if port := activeChallenge.Container.HostPort; port != 0 {
		return fmt.Sprintf("%s %s: %s (Port: %d)", activeChallenge.Container.State.ToEmoji(), title, message, port)
	} else {
		return fmt.Sprintf("%s %s: %s", activeChallenge.Container.State.ToEmoji(), title, message)
	}
}
