package repository

import (
	"context"
	"github.com/secfault-org/hacktober/internal/model/challenge"
)

type ChallengeRepository interface {
	GetAllChallenges(ctx context.Context) ([]challenge.Challenge, error)
}
