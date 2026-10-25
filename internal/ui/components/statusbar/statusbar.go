package statusbar

import (
	"time"

	"charm.land/bubbles/v2/key"
	"charm.land/bubbles/v2/spinner"
	"charm.land/bubbles/v2/textinput"
	"charm.land/bubbles/v2/timer"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/secfault-org/hacktober/internal/session"
	"github.com/secfault-org/hacktober/internal/ui/commands"
	"github.com/secfault-org/hacktober/internal/ui/styles"
	"github.com/secfault-org/hacktober/internal/ui/util"
)

type Model struct {
	screen  styles.Themed
	timeout time.Duration
	spinner spinner.Model
	errText string
	timer   timer.Model
	input   textinput.Model
}

func New(scr styles.Themed, timeout time.Duration) *Model {
	input := textinput.New()
	input.Placeholder = "Enter flag"
	input.CharLimit = 50
	return &Model{
		screen:  scr,
		timeout: timeout,
		spinner: spinner.New(spinner.WithSpinner(spinner.Jump)),
		timer:   timer.New(0),
		input:   input,
	}
}

func (bar *Model) defaultTimer() timer.Model {
	return timer.New(bar.timeout, timer.WithInterval(time.Second))
}

func (bar *Model) Init() tea.Cmd {
	return bar.spinner.Tick
}

func (bar *Model) Update(msg tea.Msg) (*Model, tea.Cmd) {
	cmds := make([]tea.Cmd, 0)
	var submitFlagKey = key.NewBinding(key.WithKeys("enter"))
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if bar.input.Focused() && key.Matches(msg, submitFlagKey) {
			submittedFlag := bar.input.Value()
			bar.input.Blur()
			bar.input.Reset()
			cmds = append(cmds, commands.FlagEntered(submittedFlag))
		}
	case commands.EnterFlagMsg:
		cmds = append(cmds, bar.input.Focus())
	case commands.EnteringFlagCanceledMsg:
		bar.input.Blur()
		bar.input.Reset()
	case commands.ChallengeStartedMsg:
		bar.errText = ""
		bar.spinner = spinner.New(spinner.WithSpinner(Clock))
		bar.timer = bar.defaultTimer()
		cmds = append(cmds, bar.spinner.Tick, bar.timer.Start())
	case commands.ChallengeStoppingMsg:
		cmds = append(cmds, bar.timer.Stop())
	case commands.ChallengeStoppedMsg:
		bar.errText = ""
		bar.timer = timer.New(0)
	case commands.ContainerErrorMsg:
		bar.errText = msg.Error()
	case timer.TimeoutMsg:
		if msg.ID == bar.timer.ID() {
			cmds = append(cmds, func() tea.Msg { return commands.ChallengeTimedOutMsg{} })
		}
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
	bar.input, cmd = bar.input.Update(msg)
	if cmd != nil {
		cmds = append(cmds, cmd)
	}
	return bar, tea.Batch(cmds...)
}

func (bar *Model) View(width int, sess *session.Session) string {
	sbStyle := bar.screen.Theme().Statusbar
	bar.spinner.Style = sbStyle.Spinner

	countdown := ""
	if bar.timer.Timeout != 0 {
		countdown = lipgloss.JoinHorizontal(lipgloss.Top,
			bar.spinner.View(),
			sbStyle.Timer.Render(bar.timer.View()),
		)
	}

	helpHint := sbStyle.Help.Render("? Help")

	w := lipgloss.Width

	main := util.ActiveChallengeStatusMessage(sess.ActiveChallenge())
	if bar.errText != "" {
		main = bar.errText
	}

	mainWidth := width - w(countdown) - w(helpHint)

	if bar.input.Focused() {
		bar.input.SetWidth(mainWidth - 4)
		main = bar.input.View()
	}

	main = sbStyle.Info.Width(mainWidth).Render(main)

	s := lipgloss.JoinHorizontal(lipgloss.Top,
		main,
		countdown,
		helpHint,
	)

	return lipgloss.NewStyle().MaxWidth(width).Render(s)
}

func (bar *Model) Height() int { return 1 }
