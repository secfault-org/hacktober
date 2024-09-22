package repository

import (
	"context"
	"github.com/secfault-org/hacktober/pkg/model"
)

type ChallengeRepository interface {
	GetAllChallenges(ctx context.Context) ([]model.Challenge, error)
}
