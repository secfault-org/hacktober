package footer

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/secfault-org/hacktober/pkg/ui/common"
)

type ToggleFooterMsg struct{}
type ToggleHelpMsg struct{}

type Footer struct {
	common common.Common
	help   help.Model
	keymap help.KeyMap
}

func New(c common.Common, keymap help.KeyMap) *Footer {
	h := help.New()
	h.Styles.ShortKey = c.Styles.HelpKey
	h.Styles.ShortDesc = c.Styles.HelpValue
	h.Styles.FullKey = c.Styles.HelpKey
	h.Styles.FullDesc = c.Styles.HelpValue

	footer := &Footer{
		common: c,
		keymap: keymap,
		help:   h,
	}
	footer.SetSize(c.Width, c.Height)
	return footer
}

func (f *Footer) SetSize(width, height int) {
	f.common.SetSize(width, height)
	f.help.Width = width - f.common.Styles.Footer.GetHorizontalFrameSize()
}

func (f *Footer) Init() tea.Cmd {
	return nil
}

func (f *Footer) Update(_ tea.Msg) (tea.Model, tea.Cmd) {
	return f, nil
}

func (f *Footer) View() string {
	if f.keymap == nil {
		return ""
	}

	style := f.common.Styles.Footer.
		Width(f.common.Width)
	helpView := f.help.View(f.keymap)
	return style.Render(helpView)
}

func (f *Footer) ShortHelp() []key.Binding {
	return f.keymap.ShortHelp()
}

func (f *Footer) FullHelp() [][]key.Binding {
	return f.keymap.FullHelp()
}

func (f *Footer) ShowAll() bool {
	return f.help.ShowAll
}

func (f *Footer) SetShowAll(show bool) {
	f.help.ShowAll = show
}

func (f *Footer) Height() int {
	return lipgloss.Height(f.View())
}

func ToggleFooterCmd() tea.Msg {
	return ToggleFooterMsg{}
}

func ToggleHelpCmd() tea.Msg {
	return ToggleHelpMsg{}
}
