package ssh

import (
	"context"

	tea "charm.land/bubbletea/v2"
	"charm.land/ssh"
	bm "charm.land/wish/v2/bubbletea"
	"github.com/secfault-org/hacktober/internal/backend"
	"github.com/secfault-org/hacktober/internal/ui"
	"github.com/secfault-org/hacktober/internal/ui/common"
)

func NewSessionHandler(ctx context.Context, b *backend.Backend) bm.Handler {
	return func(s ssh.Session) (tea.Model, []tea.ProgramOption) {
		pty, _, _ := s.Pty()

		c := common.NewCommon(ctx, pty.Window.Width, pty.Window.Height, b)

		app := ui.NewUi(c)

		return app, []tea.ProgramOption{}
	}
}
