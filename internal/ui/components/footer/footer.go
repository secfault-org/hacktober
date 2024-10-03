package footer

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/secfault-org/hacktober/internal/ui/common"
	"github.com/secfault-org/hacktober/internal/ui/components/statusbar"
)

type ToggleHelpMsg struct{}

type Footer struct {
	common    common.Common
	help      help.Model
	hideHelp  bool
	statusbar *statusbar.Model
	keymap    help.KeyMap
}

func New(c common.Common, keymap help.KeyMap) *Footer {
	h := help.New()
	h.Styles.ShortKey = c.Styles.HelpKey
	h.Styles.ShortDesc = c.Styles.HelpValue
	h.Styles.FullKey = c.Styles.HelpKey
	h.Styles.FullDesc = c.Styles.HelpValue

	footer := &Footer{
		common:    c,
		keymap:    keymap,
		help:      h,
		hideHelp:  true,
		statusbar: statusbar.New(c),
	}
	footer.SetSize(c.Width, c.Height)
	return footer
}

func (f *Footer) SetSize(width, height int) {
	f.common.SetSize(width, height)
	f.statusbar.SetSize(width, height)
	f.help.Width = width - f.common.Styles.Footer.GetHorizontalFrameSize()
}

func (f *Footer) Init() tea.Cmd {
	return nil
}

func (f *Footer) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case ToggleHelpMsg:
		f.hideHelp = !f.hideHelp
	}
	s, cmd := f.statusbar.Update(msg)
	f.statusbar = s.(*statusbar.Model)
	return f, cmd
}

func (f *Footer) View() string {
	if f.keymap == nil {
		return ""
	}

	style := f.common.Styles.Footer.
		Width(f.common.Width)
	helpView := ""
	if !f.hideHelp {
		helpView = f.help.View(f.keymap)
	}

	view := lipgloss.JoinVertical(lipgloss.Left,
		helpView,
		f.statusbar.View(),
	)
	return style.Render(view)
}

func (f *Footer) ShortHelp() []key.Binding {
	return f.keymap.ShortHelp()
}

func (f *Footer) FullHelp() [][]key.Binding {
	return f.keymap.FullHelp()
}

func (f *Footer) Height() int {
	return lipgloss.Height(f.View())
}

func ToggleHelpCmd() tea.Msg {
	return ToggleHelpMsg{}
}
