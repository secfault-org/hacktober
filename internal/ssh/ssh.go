package ssh

import (
	"context"
	"net"

	"charm.land/log/v2"
	"charm.land/ssh"
	"charm.land/wish/v2"
	"charm.land/wish/v2/activeterm"
	bm "charm.land/wish/v2/bubbletea"
	"charm.land/wish/v2/logging"
	rm "charm.land/wish/v2/recover"
	"github.com/secfault-org/hacktober/internal/application"
	"github.com/secfault-org/hacktober/internal/config"
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

func NewSSHServer(ctx context.Context, app *application.App) (*SSHServer, error) {
	cfg := app.Config()
	var err error

	s := &SSHServer{
		logger: app.Logger(),
		cfg:    cfg,
	}

	challenges, err := app.Repo().GetAllChallenges(ctx)
	if err != nil {
		return nil, err

	}

	scpHandler := NewScpChallengeHandler(challenges)

	mw := []wish.Middleware{
		rm.MiddlewareWithLogger(
			app.Logger(),
			bm.Middleware(NewSessionHandler(ctx, app)),
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
