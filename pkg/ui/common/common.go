package common

import (
	"context"
	"github.com/charmbracelet/lipgloss"
	"github.com/secfault-org/hacktober/pkg/ui/keymap"
	"github.com/secfault-org/hacktober/pkg/ui/styles"
)

type Common struct {
	ctx           context.Context
	Width, Height int
	Styles        *styles.Styles
	KeyMap        *keymap.KeyMap
	Renderer      *lipgloss.Renderer
}

func NewCommon(ctx context.Context, out *lipgloss.Renderer, width, height int) Common {
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
	}
}

func (c *Common) SetSize(width, height int) {
	c.Width = width
	c.Height = height
}
