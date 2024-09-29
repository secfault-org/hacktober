package statusbar

import (
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/secfault-org/hacktober/pkg/ui/common"
)

type Model struct {
	common      common.Common
	showSpinner bool
	spinner     spinner.Model
	info        string
}

func New(common common.Common) *Model {
	s := spinner.New(spinner.WithSpinner(spinner.Jump))
	s.Style = common.Styles.Statusbar.Spinner
	return &Model{
		common:      common,
		spinner:     s,
		showSpinner: false,
	}
}

func (bar *Model) SetSize(width, height int) {
	bar.common.SetSize(width, height)

}

func (bar *Model) Init() tea.Cmd {
	return bar.spinner.Tick
}

func (bar *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmds := make([]tea.Cmd, 0)
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		bar.SetSize(msg.Width, msg.Height)
	}
	var cmd tea.Cmd
	bar.spinner, cmd = bar.spinner.Update(msg)
	if cmd != nil {
		cmds = append(cmds, cmd)
	}
	return bar, tea.Batch(cmds...)
}

func (bar *Model) View() string {
	sbStyle := bar.common.Styles.Statusbar
	var spinView string
	if bar.showSpinner {
		spinView = bar.spinner.View()
	} else {
		spinView = ""
	}

	help := sbStyle.Help.Render("? Help")

	info := sbStyle.Info.Width(bar.common.Width - lipgloss.Width(spinView) - lipgloss.Width(help)).Render(bar.info)

	return bar.common.Renderer.NewStyle().MaxWidth(bar.common.Width).
		Render(
			lipgloss.JoinHorizontal(lipgloss.Top,
				spinView,
				info,
				help,
			),
		)
}

func (bar *Model) SetInfo(info string) {
	bar.info = info
}

func (bar *Model) SetSpinner(spinner spinner.Spinner) {
	bar.showSpinner = true
	bar.spinner.Spinner = spinner
}

func (bar *Model) HideSpinner() {
	bar.showSpinner = false
}
