package footer

import (
	"time"

	"charm.land/bubbles/v2/help"
	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/secfault-org/hacktober/internal/session"
	"github.com/secfault-org/hacktober/internal/ui/components/statusbar"
	"github.com/secfault-org/hacktober/internal/ui/styles"
)

type ToggleHelpMsg struct{}

type Footer struct {
	screen    styles.Themed
	width     int
	help      help.Model
	hideHelp  bool
	statusbar *statusbar.Model
	keymap    help.KeyMap
}

func New(scr styles.Themed, keymap help.KeyMap, timeout time.Duration) *Footer {
	return &Footer{
		screen:    scr,
		keymap:    keymap,
		help:      help.New(),
		hideHelp:  true,
		statusbar: statusbar.New(scr, timeout),
	}
}

func (f *Footer) SetWidth(w int) {
	f.width = w
	f.help.SetWidth(w - f.screen.Theme().Footer.GetHorizontalFrameSize())
}

func (f *Footer) Init() tea.Cmd {
	return nil
}

func (f *Footer) Update(msg tea.Msg) (*Footer, tea.Cmd) {
	switch msg.(type) {
	case ToggleHelpMsg:
		f.hideHelp = !f.hideHelp
	}
	s, cmd := f.statusbar.Update(msg)
	f.statusbar = s
	return f, cmd
}

func (f *Footer) View(sess *session.Session) string {
	theme := f.screen.Theme()
	inner := f.width - theme.Footer.GetHorizontalFrameSize()

	parts := make([]string, 0, 2)
	if f.showHelp() {
		parts = append(parts, f.help.View(f.keymap))
	}
	parts = append(parts, f.statusbar.View(inner, sess))

	view := lipgloss.JoinVertical(lipgloss.Left, parts...)
	return theme.Footer.Width(f.width).Render(view)
}

func (f *Footer) showHelp() bool {
	return !(f.hideHelp || f.keymap == nil)
}

func (f *Footer) ShortHelp() []key.Binding {
	return f.keymap.ShortHelp()
}

func (f *Footer) FullHelp() [][]key.Binding {
	return f.keymap.FullHelp()
}

func (f *Footer) Height() int {
	h := f.statusbar.Height()
	if f.showHelp() {
		h += lipgloss.Height(f.help.View(f.keymap))
	}

	return h + f.screen.Theme().Footer.GetVerticalFrameSize()
}

func ToggleHelpCmd() tea.Msg {
	return ToggleHelpMsg{}
}
