package challenge

import (
	"time"
)

type Challenge struct {
	Id                string
	Name              string
	Description       string
	ChallengeMarkdown string
	ReleaseDate       time.Time
	ContainerImage    string
}

func (c Challenge) Locked() bool {
	return time.Now().Before(c.ReleaseDate)
}
