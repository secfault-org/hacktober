package repository

import (
	"context"
	"github.com/secfault-org/hacktober/internal/model"
)

type ChallengeRepository interface {
	GetAllChallenges(ctx context.Context) ([]model.Challenge, error)
}
