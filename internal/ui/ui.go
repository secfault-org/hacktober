package ui

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/timer"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/secfault-org/hacktober/internal/model/challenge"
	"github.com/secfault-org/hacktober/internal/model/container"
	"github.com/secfault-org/hacktober/internal/ui/commands"
	"github.com/secfault-org/hacktober/internal/ui/common"
	"github.com/secfault-org/hacktober/internal/ui/components/confetti"
	"github.com/secfault-org/hacktober/internal/ui/components/footer"
	"github.com/secfault-org/hacktober/internal/ui/components/selector"
	"github.com/secfault-org/hacktober/internal/ui/pages/challenge_detail"
	"github.com/secfault-org/hacktober/internal/ui/pages/challenges"
	"time"
)

type page int

const (
	challengesPage page = iota
	challengeDetailsPage
)

type Ui struct {
	common       common.Common
	pages        []common.Page
	activePage   page
	footer       *footer.Footer
	confetti     tea.Model
	showConfetti bool
	inputFocused bool
	quitting     bool
}

func NewUi(c common.Common) *Ui {
	ui := &Ui{
		common:       c,
		pages:        make([]common.Page, 2),
		activePage:   challengesPage,
		showConfetti: false,
		confetti:     confetti.InitialModel(),
		inputFocused: false,
	}

	ui.footer = footer.New(c, ui)
	return ui
}

func (ui *Ui) getMargins() (wm, hm int) {
	style := ui.common.Styles.App
	wm = style.GetHorizontalFrameSize()
	hm = style.GetVerticalFrameSize()
	hm += ui.footer.Height()
	return
}

func (ui *Ui) SetSize(width, height int) {
	ui.common.SetSize(width, height)
	wm, hm := ui.getMargins()
	ui.footer.SetSize(width-wm, height-hm)
	for _, p := range ui.pages {
		if p != nil {
			p.SetSize(width-wm, height-hm)
		}
	}
}

func (ui *Ui) Init() tea.Cmd {
	ui.pages[challengesPage] = challenges.New(ui.common)
	ui.pages[challengeDetailsPage] = challenge_detail.New(ui.common)
	ui.SetSize(ui.common.Width, ui.common.Height)
	cmds := make([]tea.Cmd, 0)
	cmds = append(cmds,
		ui.pages[challengesPage].Init(),
		ui.pages[challengeDetailsPage].Init(),
	)
	return tea.Batch(cmds...)
}

func (ui *Ui) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmds := make([]tea.Cmd, 0)
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		ui.SetSize(msg.Width, msg.Height)
		for i, p := range ui.pages {
			m, cmd := p.Update(msg)
			ui.pages[i] = m.(common.Page)
			if cmd != nil {
				cmds = append(cmds, cmd)
			}
		}
	case ShowConfettiMsg:
		ui.showConfetti = true
	case HideConfettiMsg:
		ui.showConfetti = false
		cmds = append(cmds, tea.ClearScreen)
	case commands.ChallengeSolvedMsg:
		cmds = append(cmds,
			ui.confetti.Init(),
			showConfettiCmd(),
			hideConfettiCmd(2*time.Second),
		)
		cmds = append(cmds, ui.common.StopActiveChallenge()...)
	case commands.FlagEnteredMsg:
		ui.inputFocused = false
		cmds = append(cmds, commands.SubmitFlag(msg.Flag, ui.common.ActiveChallenge()))
	case tea.KeyMsg:
		if !ui.inputFocused {
			switch {
			case key.Matches(msg, ui.common.KeyMap.EnterFlag) && ui.common.IsChallengeRunning():
				ui.inputFocused = true
				cmds = append(cmds, commands.EnteringFlagCmd)
			case key.Matches(msg, ui.common.KeyMap.Quit):
				ui.quitting = true
				return ui, commands.QuittingCmd(ui.common.ActiveChallenge())
			case ui.activePage == challengeDetailsPage && key.Matches(msg, ui.common.KeyMap.Back):
				ui.activePage = challengesPage
			case key.Matches(msg, ui.common.KeyMap.Help):
				cmds = append(cmds, footer.ToggleHelpCmd)
			case key.Matches(msg, ui.common.KeyMap.StopContainer):
				cmds = append(cmds, ui.common.StopActiveChallenge()...)
			}
		} else {
			switch {
			case key.Matches(msg, ui.common.KeyMap.Back):
				ui.inputFocused = false
				cmds = append(cmds, commands.EnteringFlagCanceled)
			}
		}
	case selector.SelectMsg:
		switch msg.IdentifiableItem.(type) {
		case challenges.Item:
			chall := msg.IdentifiableItem.(challenges.Item).Challenge
			if !chall.Locked() {
				cmds = append(cmds, ui.selectChallengeCmd(chall))
			}
		}
	case commands.SelectChallengeMsg:
		ui.activePage = challengeDetailsPage
	case commands.ChallengeStaringMsg:
		activeChallenge := &challenge.ActiveChallenge{
			Challenge: msg,
			Container: &container.Container{
				State: container.Starting,
			},
		}
		cmds = append(cmds, commands.UpdateActiveChallenge(activeChallenge))
	case commands.ChallengeStartedMsg:
		ui.common.Backend.Logger().Debugf("Container for %s started", msg.Challenge.Name)
		cmds = append(cmds, commands.UpdateActiveChallenge(msg))
	case commands.ChallengeStoppingMsg:
		cmds = append(cmds, commands.UpdateActiveChallenge(msg))
	case commands.ChallengeStoppedMsg:
		ui.common.Backend.Logger().Debugf("Container for %s stopped", msg.Name)
		cmds = append(cmds, commands.UpdateActiveChallenge(nil))
		if ui.quitting {
			cmds = append(cmds, tea.Quit)
		}
	case commands.TeardownMsg:
		cmds = append(cmds, ui.common.StopActiveChallenge()...)
	case commands.ContainerErrorMsg:
		ui.common.Backend.Logger().Error(msg.Error())
	case commands.ActiveChallengeChangedMsg:
		ui.common.SetActiveChallenge(msg)
	case timer.TimeoutMsg:
		cmds = append(cmds, ui.common.StopActiveChallenge()...)
	}

	var cmd tea.Cmd
	ui.confetti, cmd = ui.confetti.Update(msg)
	if cmd != nil {
		cmds = append(cmds, cmd)
	}
	f, cmd := ui.footer.Update(msg)
	ui.footer = f.(*footer.Footer)
	if cmd != nil {
		cmds = append(cmds, cmd)
	}

	if !ui.inputFocused {
		m, cmd := ui.pages[ui.activePage].Update(msg)
		ui.pages[ui.activePage] = m.(common.Page)
		if cmd != nil {
			cmds = append(cmds, cmd)
		}
	}

	ui.SetSize(ui.common.Width, ui.common.Height)
	return ui, tea.Batch(cmds...)
}

type ShowConfettiMsg struct{}
type HideConfettiMsg struct{}

func showConfettiCmd() tea.Cmd {
	return func() tea.Msg {
		return ShowConfettiMsg{}
	}
}

func hideConfettiCmd(d time.Duration) tea.Cmd {
	return func() tea.Msg {
		time.Sleep(d)
		return HideConfettiMsg{}
	}
}

func (ui *Ui) View() string {
	if ui.showConfetti {
		return ui.confetti.View()
	}

	mainView := ui.pages[ui.activePage].View()
	if ui.quitting {
		w := lipgloss.Width(mainView)
		h := lipgloss.Height(mainView)
		message := "Quitting...\nWait until the container is stopped"
		mainView = lipgloss.PlaceVertical(h, lipgloss.Center, message)
		mainView = lipgloss.PlaceHorizontal(w, lipgloss.Center, mainView)
	}

	view := lipgloss.JoinVertical(lipgloss.Left,
		mainView,
		ui.footer.View(),
	)

	return ui.common.Styles.App.Render(view)
}

func (ui *Ui) selectChallengeCmd(challenge *challenge.Challenge) tea.Cmd {
	return func() tea.Msg {
		return commands.SelectChallengeMsg(challenge)
	}
}

func (ui *Ui) ShortHelp() []key.Binding {
	bindings := make([]key.Binding, 0)

	bindings = append(bindings, ui.pages[ui.activePage].ShortHelp()...)
	if ui.common.IsChallengeRunning() {
		bindings = append(bindings, ui.common.KeyMap.StopContainer)
	}
	bindings = append(bindings, ui.common.KeyMap.Help, ui.common.KeyMap.Quit)

	return bindings
}

func (ui *Ui) FullHelp() [][]key.Binding {
	bindings := make([][]key.Binding, 0)

	bindings = append(bindings, ui.pages[ui.activePage].FullHelp()...)
	if ui.common.IsChallengeRunning() {
		bindings = append(bindings, []key.Binding{ui.common.KeyMap.StopContainer})
	}
	bindings = append(bindings,
		[]key.Binding{
			ui.common.KeyMap.Help,
			ui.common.KeyMap.Quit,
		},
	)

	return bindings
}
