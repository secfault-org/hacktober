package repository

import "context"

type ContextKey int

const (
	repositoryKey ContextKey = iota
)

func FromContext(ctx context.Context) Repository {
	if repo, ok := ctx.Value(repositoryKey).(Repository); ok {
		return repo
	}
	return nil
}

func WithContext(ctx context.Context, repo Repository) context.Context {
	return context.WithValue(ctx, repositoryKey, repo)
}
