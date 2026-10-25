package ui

import (
	"context"
	"time"

	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/secfault-org/hacktober/internal/application"
	"github.com/secfault-org/hacktober/internal/model/challenge"
	"github.com/secfault-org/hacktober/internal/session"
	"github.com/secfault-org/hacktober/internal/ui/commands"
	"github.com/secfault-org/hacktober/internal/ui/common"
	"github.com/secfault-org/hacktober/internal/ui/components/confetti"
	"github.com/secfault-org/hacktober/internal/ui/components/footer"
	"github.com/secfault-org/hacktober/internal/ui/components/selector"
	"github.com/secfault-org/hacktober/internal/ui/keymap"
	"github.com/secfault-org/hacktober/internal/ui/pages/challenge_detail"
	"github.com/secfault-org/hacktober/internal/ui/pages/challenges"
	"github.com/secfault-org/hacktober/internal/ui/screen"
)

type page int

const (
	challengesPage page = iota
	challengeDetailsPage
)

type Ui struct {
	ctx          context.Context
	app          *application.App
	session      *session.Session
	keyMap       *keymap.KeyMap
	screen       *screen.Screen
	pages        []common.Page
	activePage   page
	footer       *footer.Footer
	confetti     confetti.Model
	showConfetti bool
	inputFocused bool
}

func New(ctx context.Context, a *application.App, sess *session.Session, w, h int) *Ui {
	scr := screen.New(w, h)
	keyMap := keymap.DefaultKeyMap()

	timeout := time.Duration(a.Config().Container.TimeoutInMinutes) * time.Minute

	ui := &Ui{
		ctx:          ctx,
		app:          a,
		session:      sess,
		screen:       scr,
		keyMap:       keyMap,
		pages:        make([]common.Page, 2),
		activePage:   challengesPage,
		showConfetti: false,
		confetti:     confetti.InitialModel(scr),
	}
	ui.footer = footer.New(scr, ui, timeout)
	return ui
}

func (ui *Ui) SetSize(width, height int) {
	ui.screen.Width, ui.screen.Height = width, height
}

func (ui *Ui) layout() {
	appStyle := ui.screen.Theme().App
	usableWidth := ui.screen.Width - appStyle.GetHorizontalFrameSize()
	usableHeight := ui.screen.Height - appStyle.GetVerticalFrameSize()

	ui.footer.SetWidth(usableWidth)
	mainHeight := usableHeight - ui.footer.Height()

	for _, page := range ui.pages {
		if page != nil {
			page.SetSize(usableWidth, mainHeight)
		}
	}
}

func (ui *Ui) Init() tea.Cmd {
	challs, err := ui.app.Repo().GetAllChallenges(ui.ctx)

	if err != nil {
		return common.ErrorCmd(err)
	}

	ui.pages[challengesPage] = challenges.New(ui.screen, ui.keyMap, challs)
	ui.pages[challengeDetailsPage] = challenge_detail.New(ui.screen, ui.keyMap, ui.session, ui.ctx)
	cmds := make([]tea.Cmd, 0)
	cmds = append(cmds,
		ui.pages[challengesPage].Init(),
		ui.pages[challengeDetailsPage].Init(),
		ui.confetti.Init(),
	)

	ui.layout()

	return tea.Batch(cmds...)
}

func (ui *Ui) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmds := make([]tea.Cmd, 0)
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		ui.screen.Width, ui.screen.Height = msg.Width, msg.Height
	case ShowConfettiMsg:
		ui.confetti.Reset()
		ui.showConfetti = true
	case HideConfettiMsg:
		ui.showConfetti = false
		cmds = append(cmds, tea.ClearScreen)
	case commands.ChallengeSolvedMsg:
		cmds = append(cmds,
			showConfettiCmd(),
			hideConfettiCmd(2*time.Second),
			commands.StopChallenge(ui.session, ui.ctx),
		)
	case commands.FlagEnteredMsg:
		ui.inputFocused = false
		cmds = append(cmds, commands.SubmitFlag(ui.session, msg.Flag))
	case tea.KeyMsg:
		if !ui.inputFocused {
			switch {
			case key.Matches(msg, ui.keyMap.EnterFlag) && ui.session.IsChallengeRunning():
				ui.inputFocused = true
				cmds = append(cmds, commands.EnteringFlagCmd)
			case key.Matches(msg, ui.keyMap.Quit):
				return ui, tea.Quit
			case ui.activePage == challengeDetailsPage && key.Matches(msg, ui.keyMap.Back):
				ui.activePage = challengesPage
			case key.Matches(msg, ui.keyMap.Help):
				cmds = append(cmds, footer.ToggleHelpCmd)
			case key.Matches(msg, ui.keyMap.StopContainer):
				cmds = append(cmds, commands.StopChallenge(ui.session, ui.ctx))
			}
		} else {
			switch {
			case key.Matches(msg, ui.keyMap.Back):
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
		cmds = append(cmds, commands.SpawnContainer(ui.session, ui.ctx, msg))
	case commands.ChallengeStartedMsg:
		ui.app.Logger().Debugf("Container for %s started", msg.Challenge.Name)
	case commands.ChallengeStoppedMsg:
		ui.app.Logger().Debugf("Container for %s stopped", msg.Name)
	case commands.ContainerErrorMsg:
		ui.app.Logger().Error(msg.Error())
	case commands.ChallengeTimedOutMsg:
		if ui.session.CanStopChallenge() {
			cmds = append(cmds, commands.StopChallenge(ui.session, ui.ctx))
		}
	}

	var cmd tea.Cmd
	con, cmd := ui.confetti.Update(msg)
	ui.confetti = con
	if cmd != nil {
		cmds = append(cmds, cmd)
	}
	f, cmd := ui.footer.Update(msg)
	ui.footer = f
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

	switch msg.(type) {
	case tea.WindowSizeMsg, footer.ToggleHelpMsg:
		ui.layout()
	}
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

func (ui *Ui) content() string {
	if ui.showConfetti {
		return ui.confetti.View()
	}

	appStyle := ui.screen.Theme().App
	usableHeight := ui.screen.Height - appStyle.GetVerticalFrameSize()

	mainHeight := usableHeight - ui.footer.Height()

	mainView := lipgloss.NewStyle().
		Height(mainHeight).
		MaxHeight(mainHeight).
		Render(ui.pages[ui.activePage].View())

	view := lipgloss.JoinVertical(lipgloss.Left,
		mainView,
		ui.footer.View(ui.session),
	)

	return ui.screen.Theme().App.Render(view)
}

func (ui *Ui) View() tea.View {
	view := tea.NewView(ui.content())
	view.AltScreen = true
	return view
}

func (ui *Ui) selectChallengeCmd(challenge *challenge.Challenge) tea.Cmd {
	return func() tea.Msg {
		return commands.SelectChallengeMsg(challenge)
	}
}

func (ui *Ui) ShortHelp() []key.Binding {
	bindings := make([]key.Binding, 0)

	bindings = append(bindings, ui.pages[ui.activePage].ShortHelp()...)
	if ui.session.CanStopChallenge() {
		bindings = append(bindings, ui.keyMap.StopContainer)
	}
	bindings = append(bindings, ui.keyMap.Help, ui.keyMap.Quit)

	return bindings
}

func (ui *Ui) FullHelp() [][]key.Binding {
	bindings := make([][]key.Binding, 0)

	bindings = append(bindings, ui.pages[ui.activePage].FullHelp()...)
	if ui.session.CanStopChallenge() {
		bindings = append(bindings, []key.Binding{ui.keyMap.StopContainer})
	}
	bindings = append(bindings,
		[]key.Binding{
			ui.keyMap.Help,
			ui.keyMap.Quit,
		},
	)

	return bindings
}
