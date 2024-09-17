package challenge_detail

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"hacktopber2024/pkg/model"
	"hacktopber2024/pkg/ui/common"
)

type SelectChallengeMsg model.Challenge

type GoBackMsg struct{}

type ChallengeDetailPage struct {
	common            common.Common
	selectedChallenge model.Challenge
}

func New(common common.Common) *ChallengeDetailPage {
	return &ChallengeDetailPage{
		common: common,
	}
}

func (c *ChallengeDetailPage) getMargins() (int, int) {
	hm := c.common.Styles.ChallengeDetail.Body.GetVerticalFrameSize()
	return 0, hm
}

func (c *ChallengeDetailPage) SetSize(width, height int) {
	c.common.SetSize(width, height)
}

func (c *ChallengeDetailPage) commonHelp() []key.Binding {
	b := make([]key.Binding, 0)
	back := c.common.KeyMap.Back
	back.SetHelp("esc", "back to challenge list")
	b = append(b, back)
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
	return nil
}

func (c *ChallengeDetailPage) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmds := make([]tea.Cmd, 0)
	switch msg := msg.(type) {
	case SelectChallengeMsg:
		c.selectedChallenge = model.Challenge(msg)
		cmds = append(cmds,
			c.Init(),
		)
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, c.common.KeyMap.Back):
			cmds = append(cmds, goBackCmd)
		}
	case tea.WindowSizeMsg:
		c.SetSize(msg.Width, msg.Height)
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
	main := c.selectedChallenge.Description
	view := mainStyle.Render(main)
	return s.Render(view)
}

func goBackCmd() tea.Msg {
	return GoBackMsg{}
}
