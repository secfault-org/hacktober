package backend

import "context"

type ContextKey int

const (
	backendKey ContextKey = iota
)

func FromContext(ctx context.Context) *Backend {
	if backend, ok := ctx.Value(backendKey).(*Backend); ok {
		return backend
	}
	return nil
}

func WithContext(ctx context.Context, backend *Backend) context.Context {
	return context.WithValue(ctx, backendKey, backend)
}
