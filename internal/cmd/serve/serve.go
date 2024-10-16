package serve

import (
	"context"
	"errors"
	"github.com/charmbracelet/log"
	cssh "github.com/charmbracelet/ssh"
	"github.com/secfault-org/hacktober/internal/backend"
	"github.com/secfault-org/hacktober/internal/config"
	"github.com/secfault-org/hacktober/internal/container/podman"
	"github.com/secfault-org/hacktober/internal/repository"
	"github.com/secfault-org/hacktober/internal/ssh"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	Command = &cobra.Command{
		Use:   "serve",
		Short: "Start the Hacktober server",
		Args:  cobra.NoArgs,
		RunE: func(c *cobra.Command, _ []string) error {
			cfg := config.NewConfig(viper.GetViper())
			ctx := c.Context()
			logger := log.FromContext(ctx).WithPrefix("backend")
			l := logger.WithPrefix("serve")

			containerService := podman.NewContainerService(ctx)
			ctx, err := containerService.Connect(ctx)
			if err != nil {
				l.Error("Could not connect to Podman", "error", err)
				return err
			}

			b := backend.NewBackend(cfg, repository.NewRepository(ctx, cfg.Challenge.BaseDir), logger, containerService)

			server, err := ssh.NewSSHServer(ctx, b)
			if err != nil {
				l.Error("Could not start SSH Server", "error", err)
				return err
			}

			done := make(chan os.Signal, 1)
			signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

			go func() {
				if err := server.Start(); err != nil && !errors.Is(err, cssh.ErrServerClosed) {
					l.Error("Error starting Server", "error", err)
					done <- nil
				}
			}()

			<-done
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer func() { cancel() }()
			if err := server.Shutdown(ctx); err != nil && !errors.Is(err, cssh.ErrServerClosed) {
				log.Error("Error stopping Server", "error", err)
			}

			return nil
		},
	}
)
