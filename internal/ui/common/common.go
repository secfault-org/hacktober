package common

import (
	"context"
	"github.com/charmbracelet/lipgloss"
	"github.com/secfault-org/hacktober/internal/backend"
	"github.com/secfault-org/hacktober/internal/container"
	"github.com/secfault-org/hacktober/internal/repository"
	"github.com/secfault-org/hacktober/internal/ui/keymap"
	"github.com/secfault-org/hacktober/internal/ui/styles"
)

type Common struct {
	ctx           context.Context
	Width, Height int
	Styles        *styles.Styles
	KeyMap        *keymap.KeyMap
	Renderer      *lipgloss.Renderer
	Backend       *backend.Backend
}

func NewCommon(ctx context.Context, out *lipgloss.Renderer, width, height int, b *backend.Backend) Common {
	if ctx == nil {
		ctx = context.TODO()
	}
	return Common{
		ctx:      ctx,
		Width:    width,
		Height:   height,
		Renderer: out,
		Styles:   styles.DefaultStyles(out),
		KeyMap:   keymap.DefaultKeyMap(),
		Backend:  b,
	}
}

func (c *Common) SetSize(width, height int) {
	c.Width = width
	c.Height = height
}

func (c *Common) Repo() repository.Repository {
	return c.Backend.Repo
}

func (c *Common) ContainerService() container.Service {
	return c.Backend.ContainerService
}

func (c *Common) Context() context.Context {
	return c.ctx
}
