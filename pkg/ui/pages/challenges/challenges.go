package challenges

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/secfault-org/hacktober/pkg/ui/common"
	"github.com/secfault-org/hacktober/pkg/ui/components/selector"
)

type ChallengePage struct {
	common   common.Common
	selector *selector.Selector
}

func New(common common.Common) *ChallengePage {

	list := selector.New(
		common,
		[]selector.IdentifiableItem{},
		NewItemDelegate(&common),
	)

	list.Title = "Challenges - Hacktober 2024"
	list.SetShowTitle(true)
	list.SetShowHelp(false)
	list.SetShowStatusBar(false)
	list.DisableQuitKeybindings()

	return &ChallengePage{
		common:   common,
		selector: list,
	}
}

func (c *ChallengePage) Init() tea.Cmd {
	ctx := c.common.Context()
	challenges, err := c.common.Repo().GetAllChallenges(ctx)
	if err != nil {
		return common.ErrorCmd(err)
	}

	items := make([]selector.IdentifiableItem, len(challenges))
	for i, it := range challenges {
		items[i] = Item{Challenge: it}
	}

	return tea.Batch(
		c.selector.Init(),
		c.selector.SetItems(items),
	)
}

func (c *ChallengePage) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmds := make([]tea.Cmd, 0)

	m, cmd := c.selector.Update(msg)
	c.selector = m.(*selector.Selector)
	if cmd != nil {
		cmds = append(cmds, cmd)
	}

	return c, tea.Batch(cmds...)
}

func (c *ChallengePage) View() string {
	ss := c.common.Renderer.NewStyle().
		Width(c.common.Width).
		Height(c.common.Height)
	view := ss.Render(c.selector.View())
	return lipgloss.JoinVertical(lipgloss.Left, view)
}

func (c *ChallengePage) getMargins() (wm, hm int) {
	wm = 0
	hm = 0
	return
}

func (c *ChallengePage) SetSize(width, height int) {
	c.common.SetSize(width, height)
	wm, hm := c.getMargins()
	c.selector.SetSize(width-wm, height-hm)
}

func (c *ChallengePage) ShortHelp() []key.Binding {
	keyBindings := make([]key.Binding, 0)
	keyBindings = append(keyBindings,
		c.common.KeyMap.UpDown,
		c.common.KeyMap.Select,
	)

	return keyBindings
}

func (c *ChallengePage) FullHelp() [][]key.Binding {
	bindings := [][]key.Binding{}
	bindings = append(bindings, []key.Binding{
		c.selector.KeyMap.CursorUp,
		c.selector.KeyMap.CursorDown,
	})

	return bindings
}
