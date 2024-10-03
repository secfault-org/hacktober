package statusbar

import (
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/timer"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/secfault-org/hacktober/internal/ui/commands"
	"github.com/secfault-org/hacktober/internal/ui/common"
	"github.com/secfault-org/hacktober/internal/ui/util"
	"time"
)

type Model struct {
	common  common.Common
	spinner spinner.Model
	info    string
	timer   timer.Model
}

func New(common common.Common) *Model {
	s := spinner.New(spinner.WithSpinner(spinner.Jump))
	s.Style = common.Styles.Statusbar.Spinner
	return &Model{
		common:  common,
		spinner: s,
		timer:   timer.New(0),
	}
}

func defaultTimer() timer.Model {
	return timer.NewWithInterval(10*time.Minute, time.Second)
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
	case commands.ActiveChallengeChangedMsg:
		bar.info = util.ActiveChallengeStatusMessage(msg)
	case commands.ChallengeStartedMsg:
		bar.spinner = spinner.New(spinner.WithSpinner(Clock))
		bar.timer = defaultTimer()
		cmds = append(cmds, bar.spinner.Tick, bar.timer.Start())
	case commands.ChallengeStoppingMsg:
		cmds = append(cmds, bar.timer.Stop())
	case commands.ChallengeStoppedMsg:
		bar.timer = timer.New(0)
	case commands.ContainerErrorMsg:
		bar.info = msg.Error()
	}
	var cmd tea.Cmd
	bar.spinner, cmd = bar.spinner.Update(msg)
	if cmd != nil {
		cmds = append(cmds, cmd)
	}
	bar.timer, cmd = bar.timer.Update(msg)
	if cmd != nil {
		cmds = append(cmds, cmd)
	}
	return bar, tea.Batch(cmds...)
}

func (bar *Model) View() string {
	sbStyle := bar.common.Styles.Statusbar

	countdown := ""
	if bar.timer.Timeout != 0 {
		countdown = sbStyle.Timer.Render(
			lipgloss.JoinHorizontal(lipgloss.Top,
				bar.spinner.View(),
				bar.timer.View(),
			),
		)
	}

	help := sbStyle.Help.Render("? Help")

	w := lipgloss.Width

	info := sbStyle.Info.Width(bar.common.Width - w(countdown) - w(help)).Render(bar.info)

	return bar.common.Renderer.NewStyle().MaxWidth(bar.common.Width).
		Render(
			lipgloss.JoinHorizontal(lipgloss.Top,
				info,
				countdown,
				help,
			),
		)
}
