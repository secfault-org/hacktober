package main

import (
	"context"
	"errors"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish/activeterm"
	bm "github.com/charmbracelet/wish/bubbletea"
	"github.com/charmbracelet/wish/logging"
	"github.com/secfault-org/hacktober/pkg/backend"
	"github.com/secfault-org/hacktober/pkg/container/podman"
	"github.com/secfault-org/hacktober/pkg/repository"
	"github.com/secfault-org/hacktober/pkg/ui"
	"github.com/secfault-org/hacktober/pkg/ui/common"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/charmbracelet/wish"
)

const (
	host = "localhost"
	port = "22222"
)

func main() {
	s, err := wish.NewServer(
		wish.WithAddress(net.JoinHostPort(host, port)),
		wish.WithHostKeyPath(".ssh/id_ed25519"),
		wish.WithPublicKeyAuth(func(ctx ssh.Context, key ssh.PublicKey) bool {
			return true
		}),
		wish.WithMiddleware(
			bm.Middleware(teaHandler),
			activeterm.Middleware(),
			logging.Middleware(),
		),
	)

	if err != nil {
		log.Error("Could not start Server", "error", err)
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	log.Info("Starting SSH Server", "host", host, "port", port)
	go func() {
		if err := s.ListenAndServe(); err != nil && !errors.Is(err, ssh.ErrServerClosed) {
			log.Error("Error starting Server", "error", err)
			done <- nil
		}
	}()

	<-done
	log.Info("Stopping SSH Server")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() { cancel() }()
	if err := s.Shutdown(ctx); err != nil && !errors.Is(err, ssh.ErrServerClosed) {
		log.Error("Error stopping Server", "error", err)
	}
}

func teaHandler(s ssh.Session) (tea.Model, []tea.ProgramOption) {
	pty, _, _ := s.Pty()
	ctx := context.Background()
	containerService := podman.NewContainerService(ctx)
	ctx = backend.WithContext(ctx, backend.NewBackend(ctx, repository.NewRepository(ctx, "challenges/2024"), containerService))

	renderer := bm.MakeRenderer(s)
	c := common.NewCommon(ctx, renderer, pty.Window.Width, pty.Window.Height)

	app := ui.NewUi(c)

	return app, []tea.ProgramOption{tea.WithAltScreen()}
}
