package common

import (
	"context"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/secfault-org/hacktober/internal/backend"
	"github.com/secfault-org/hacktober/internal/container"
	"github.com/secfault-org/hacktober/internal/model/challenge"
	cm "github.com/secfault-org/hacktober/internal/model/container"
	"github.com/secfault-org/hacktober/internal/repository"
	"github.com/secfault-org/hacktober/internal/ui/commands"
	"github.com/secfault-org/hacktober/internal/ui/keymap"
	"github.com/secfault-org/hacktober/internal/ui/styles"
)

type Common struct {
	ctx             context.Context
	Width, Height   int
	Styles          *styles.Styles
	KeyMap          *keymap.KeyMap
	Renderer        *lipgloss.Renderer
	Backend         *backend.Backend
	activeChallenge *challenge.ActiveChallenge
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

func (c *Common) IsChallengeRunning() bool {
	return c.activeChallenge != nil && c.activeChallenge.IsRunning()
}

func (c *Common) SetActiveChallenge(challenge *challenge.ActiveChallenge) {
	c.activeChallenge = challenge
}

func (c *Common) ActiveChallenge() *challenge.ActiveChallenge {
	return c.activeChallenge
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

func (c *Common) StopActiveChallenge() []tea.Cmd {
	if c.activeChallenge == nil {
		return nil
	}

	c.activeChallenge.Container.State = cm.Stopping

	stopChallengeCmd := func() tea.Msg {
		c.KeyMap.StopContainer.SetEnabled(false)
		err := c.ContainerService().StopContainer(c.ctx, c.activeChallenge.Container.ID)
		if err != nil {
			c.KeyMap.StopContainer.SetEnabled(true)
			return commands.ContainerErrorMsg(err)
		}
		c.KeyMap.SpawnContainer.SetEnabled(true)
		chall := c.activeChallenge.Challenge
		c.activeChallenge = nil
		return commands.ChallengeStoppedMsg(chall)
	}

	return []tea.Cmd{
		commands.ChallengeStopping(c.activeChallenge),
		stopChallengeCmd,
	}
}

func (c *Common) StartChallenge(chall *challenge.Challenge) []tea.Cmd {
	c.KeyMap.SpawnContainer.SetEnabled(false)

	c.activeChallenge = &challenge.ActiveChallenge{
		Challenge: chall,
		Container: &cm.Container{
			State: cm.Starting,
		},
	}

	spawnCmd := func() tea.Msg {
		runningContainer, err := c.ContainerService().StartContainer(c.ctx, chall.ContainerImage, 1337)
		if err != nil {
			c.KeyMap.SpawnContainer.SetEnabled(true)
			return commands.ContainerErrorMsg(err)
		}
		c.KeyMap.StopContainer.SetEnabled(true)
		c.activeChallenge = &challenge.ActiveChallenge{
			Challenge: chall,
			Container: runningContainer,
		}
		return commands.ChallengeStartedMsg(c.activeChallenge)
	}

	return []tea.Cmd{
		commands.ChallengeStarting(chall),
		spawnCmd,
	}
}
