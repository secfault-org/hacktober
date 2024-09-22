package repository

import (
	"context"
)

type Repository interface {
	ChallengeRepository
}

type datastore struct {
	ctx context.Context

	*challengeRepo
}

func NewRepository(ctx context.Context, basedir string) Repository {
	return &datastore{
		ctx: ctx,
		challengeRepo: &challengeRepo{
			basedir: basedir,
		},
	}
}
