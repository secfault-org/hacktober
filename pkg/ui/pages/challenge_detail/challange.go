package challenge_detail

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
	"github.com/secfault-org/hacktober/pkg/model"
	"github.com/secfault-org/hacktober/pkg/ui/commands"
	"github.com/secfault-org/hacktober/pkg/ui/common"
	"github.com/secfault-org/hacktober/pkg/ui/components/statusbar"
	"github.com/secfault-org/hacktober/pkg/ui/components/viewport"
	"time"
)

const (
	ContainerStateStarting = iota
	ContainerStateRunning
	ContainerStateStopped
)

type ChallengeDetailPage struct {
	*viewport.Viewport
	common            common.Common
	selectedChallenge model.Challenge
	ContainerState    int
	statusbar         *statusbar.Model
}

func New(common common.Common) *ChallengeDetailPage {
	page := &ChallengeDetailPage{
		common:         common,
		Viewport:       viewport.New(common),
		ContainerState: ContainerStateStopped,
		statusbar:      statusbar.New(common),
	}
	page.SetSize(common.Height, common.Width)
	return page
}

func (c *ChallengeDetailPage) getMargins() (int, int) {
	hm := c.common.Styles.ChallengeDetail.Body.GetVerticalFrameSize() +
		c.common.Styles.Statusbar.Base.GetHeight()
	return 0, hm
}

func (c *ChallengeDetailPage) SetSize(width, height int) {
	c.common.SetSize(width, height)
	_, hm := c.getMargins()
	c.Viewport.SetSize(width, height-hm)
	c.statusbar.SetSize(width, height-hm)
}

func (c *ChallengeDetailPage) commonHelp() []key.Binding {
	b := make([]key.Binding, 0)
	back := c.common.KeyMap.Back
	back.SetHelp("esc", "back to challenge list")
	var containerHelp key.Binding
	if c.ContainerState == ContainerStateRunning {
		containerHelp = c.common.KeyMap.StopContainer
	} else {
		containerHelp = c.common.KeyMap.SpawnContainer
	}
	b = append(b,
		back,
		containerHelp,
	)
	return b
}

func (c *ChallengeDetailPage) ShortHelp() []key.Binding {
	return c.commonHelp()
}

func (c *ChallengeDetailPage) FullHelp() [][]key.Binding {
	b := make([][]key.Binding, 0)
	b = append(b, c.commonHelp())
	return b
}

func (c *ChallengeDetailPage) Init() tea.Cmd {
	if c.selectedChallenge.ChallengeMarkdown == "" {
		c.Viewport.Model.SetContent("Loading...")
		return nil
	} else {
		width := c.common.Width
		if width > 120 {
			width = 120
		}
		r, err := glamour.NewTermRenderer(
			glamour.WithAutoStyle(),
			glamour.WithWordWrap(width),
		)
		if err != nil {
			return nil
		}
		rendered, err := r.Render(c.selectedChallenge.ChallengeMarkdown)
		if err != nil {
			return nil
		}
		c.Viewport.Model.SetContent(rendered)
	}

	return nil
}

func (c *ChallengeDetailPage) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmds := make([]tea.Cmd, 0)
	switch msg := msg.(type) {
	case commands.SelectChallengeMsg:
		c.selectedChallenge = model.Challenge(msg)
		cmds = append(cmds,
			c.Init(),
		)
	case commands.ContainerSpawnedMsg:
		c.ContainerState = msg.State
		cmds = append(cmds, spawnContainerLoading(false, ""))
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, c.common.KeyMap.Back):
			cmds = append(cmds, goBackCmd)
		case key.Matches(msg, c.common.KeyMap.SpawnContainer) && c.ContainerState == ContainerStateStopped:
			cmds = append(cmds,
				spawnContainerLoading(true, "Spawning container..."),
				spawnContainerCmd(c.selectedChallenge),
			)
		}
	case tea.WindowSizeMsg:
		c.SetSize(msg.Width, msg.Height)
	}
	s, cmd := c.statusbar.Update(msg)
	c.statusbar = s.(*statusbar.Model)
	if cmd != nil {
		cmds = append(cmds, cmd)
	}

	v, cmd := c.Viewport.Update(msg)
	c.Viewport = v.(*viewport.Viewport)
	if cmd != nil {
		cmds = append(cmds, cmd)
	}
	return c, tea.Batch(cmds...)
}

func (c *ChallengeDetailPage) View() string {
	wm, hm := c.getMargins()
	s := c.common.Styles.ChallengeDetail.Base.
		Width(c.common.Width - wm).
		Height(c.common.Height - hm)
	mainStyle := c.common.Styles.ChallengeDetail.Body.
		Height(c.common.Height - hm)
	main := c.Viewport.View()
	view := lipgloss.JoinVertical(lipgloss.Left,
		mainStyle.Render(main),
		c.statusbar.View(),
	)
	return s.Render(view)
}

func goBackCmd() tea.Msg {
	return commands.GoBackMsg{}
}

func spawnContainerLoading(loading bool, msg string) tea.Cmd {
	return func() tea.Msg {
		return commands.ContainerLoadingMsg{Loading: loading, Message: msg}
	}
}

func spawnContainerCmd(challenge model.Challenge) tea.Cmd {
	return func() tea.Msg {
		time.Sleep(2 * time.Second)
		return commands.ContainerSpawnedMsg{
			Challenge: challenge,
			State:     ContainerStateRunning,
		}
	}
}
