package ssh

import (
	"context"
	"github.com/charmbracelet/log"
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	"github.com/charmbracelet/wish/activeterm"
	bm "github.com/charmbracelet/wish/bubbletea"
	"github.com/charmbracelet/wish/logging"
	rm "github.com/charmbracelet/wish/recover"
	"github.com/secfault-org/hacktober/internal/backend"
	"github.com/secfault-org/hacktober/internal/config"
	"net"
)

type SSHServer struct {
	logger *log.Logger
	srv    *ssh.Server
	cfg    *config.Config
}

func (s *SSHServer) PublicKeyHandler(ctx ssh.Context, pk ssh.PublicKey) bool {
	if pk == nil {
		return false
	}

	//user, _ := s.be.UserByPublicKey(ctx, pk)
	//if user != nil {
	//	ctx.SetValue(proto.ContextKeyUser, user)
	//}

	return true
}

func NewSSHServer(ctx context.Context, backend *backend.Backend) (*SSHServer, error) {
	cfg := backend.Config
	var err error

	s := &SSHServer{
		logger: backend.Logger(),
		cfg:    cfg,
	}

	challenges, err := backend.Repo.GetAllChallenges(ctx)
	if err != nil {
		return nil, err

	}

	scpHandler := NewScpChallengeHandler(challenges)

	mw := []wish.Middleware{
		rm.MiddlewareWithLogger(
			backend.Logger(),
			bm.Middleware(NewSessionHandler(ctx, backend)),
			activeterm.Middleware(),
			logging.Middleware(),
		),
	}

	opts := []ssh.Option{
		wish.WithAddress(net.JoinHostPort(cfg.SSH.Host, cfg.SSH.Port)),
		wish.WithHostKeyPath(cfg.SSH.KeyPath),
		wish.WithPublicKeyAuth(s.PublicKeyHandler),
		wish.WithSubsystem("sftp", s.sftpSubsystem(scpHandler)),
		wish.WithMiddleware(mw...),
		ssh.AllocatePty(),
	}

	s.srv, err = wish.NewServer(opts...)

	if err != nil {
		return nil, err
	}

	return s, nil
}

func (s *SSHServer) Start() error {
	s.logger.Info("Starting SSH Server", "host", s.cfg.SSH.Host, "port", s.cfg.SSH.Port)
	return s.srv.ListenAndServe()
}

func (s *SSHServer) Shutdown(ctx context.Context) error {
	s.logger.Info("Stopping SSH Server")
	return s.srv.Shutdown(ctx)
}
