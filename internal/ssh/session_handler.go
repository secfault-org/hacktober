package ssh

import (
	"context"
	"time"

	tea "charm.land/bubbletea/v2"
	"charm.land/ssh"
	bm "charm.land/wish/v2/bubbletea"
	"github.com/secfault-org/hacktober/internal/application"
	"github.com/secfault-org/hacktober/internal/session"
	"github.com/secfault-org/hacktober/internal/ui"
)

func NewSessionHandler(ctx context.Context, app *application.App) bm.Handler {
	return func(s ssh.Session) (tea.Model, []tea.ProgramOption) {
		pty, _, _ := s.Pty()

		sess := session.New(app.ContainerService(), app.Config())

		go func() {
			<-s.Context().Done()
			stopCtx, cancel := context.WithTimeout(ctx, 20*time.Second)
			defer cancel()
			if err := sess.Stop(stopCtx); err != nil {
				app.Logger().Error(err)
			}
		}()

		uiView := ui.New(ctx, app, sess, pty.Window.Width, pty.Window.Height)

		return uiView, []tea.ProgramOption{}
	}
}
