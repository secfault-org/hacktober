package challenge_detail

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/secfault-org/hacktober/internal/model/challenge"
	"github.com/secfault-org/hacktober/internal/ui/commands"
	"github.com/secfault-org/hacktober/internal/ui/common"
	"github.com/secfault-org/hacktober/internal/ui/components/viewport"
)

type ChallengeDetailPage struct {
	viewport          *viewport.Viewport
	common            common.Common
	selectedChallenge *challenge.Challenge
}

func New(common common.Common) *ChallengeDetailPage {
	page := &ChallengeDetailPage{
		common:   common,
		viewport: viewport.New(common),
	}
	page.SetSize(common.Height, common.Width)
	return page
}

func (c *ChallengeDetailPage) Init() tea.Cmd {
	if c.selectedChallenge == nil {
		c.viewport.Model.SetContent("Loading...")
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
		c.viewport.Model.SetContent(rendered)
	}
	return nil
}

func (c *ChallengeDetailPage) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmds := make([]tea.Cmd, 0)
	switch msg := msg.(type) {
	case commands.SelectChallengeMsg:
		c.selectedChallenge = msg
		cmds = append(cmds,
			c.Init(),
		)
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, c.common.KeyMap.Back):
			cmds = append(cmds, goBackCmd)
		case key.Matches(msg, c.common.KeyMap.SpawnContainer) && c.selectedChallenge != nil:
			cmds = append(cmds, c.common.StartChallenge(c.selectedChallenge)...)
		}
	case tea.WindowSizeMsg:
		c.SetSize(msg.Width, msg.Height)
	}
	v, cmd := c.viewport.Update(msg)
	c.viewport = v.(*viewport.Viewport)
	if cmd != nil {
		cmds = append(cmds, cmd)
	}
	return c, tea.Batch(cmds...)
}

func (c *ChallengeDetailPage) View() string {
	_, hm := c.getMargins()
	mainStyle := c.common.Styles.ChallengeDetail.Body.
		Height(c.common.Height - hm)

	return mainStyle.Render(c.viewport.View())
}

func (c *ChallengeDetailPage) getMargins() (int, int) {
	hm := c.common.Styles.ChallengeDetail.Body.GetVerticalFrameSize()
	return 0, hm
}

func (c *ChallengeDetailPage) SetSize(width, height int) {
	c.common.SetSize(width, height)
	_, hm := c.getMargins()
	c.viewport.SetSize(width, height-hm)
}

func (c *ChallengeDetailPage) commonHelp() []key.Binding {
	b := make([]key.Binding, 0)
	back := c.common.KeyMap.Back
	back.SetHelp("esc", "back to challenge list")
	if !c.common.IsChallengeRunning() {
		b = append(b, c.common.KeyMap.SpawnContainer)
	}
	b = append(b,
		back,
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

func goBackCmd() tea.Msg {
	return commands.GoBackMsg{}
}
