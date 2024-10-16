package ssh

import (
	"context"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/ssh"
	bm "github.com/charmbracelet/wish/bubbletea"
	"github.com/secfault-org/hacktober/internal/backend"
	"github.com/secfault-org/hacktober/internal/ui"
	"github.com/secfault-org/hacktober/internal/ui/common"
)

func NewSessionHandler(ctx context.Context, b *backend.Backend) bm.Handler {
	return func(s ssh.Session) (tea.Model, []tea.ProgramOption) {
		pty, _, _ := s.Pty()

		renderer := bm.MakeRenderer(s)
		c := common.NewCommon(ctx, renderer, pty.Window.Width, pty.Window.Height, b)

		app := ui.NewUi(c)

		return app, []tea.ProgramOption{tea.WithAltScreen()}
	}
}
