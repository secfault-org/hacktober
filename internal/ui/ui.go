package ui

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/secfault-org/hacktober/internal/model/challenge"
	"github.com/secfault-org/hacktober/internal/ui/commands"
	"github.com/secfault-org/hacktober/internal/ui/common"
	"github.com/secfault-org/hacktober/internal/ui/components/footer"
	"github.com/secfault-org/hacktober/internal/ui/components/selector"
	"github.com/secfault-org/hacktober/internal/ui/pages/challenge_detail"
	"github.com/secfault-org/hacktober/internal/ui/pages/challenges"
)

type page int

const (
	challengesPage page = iota
	challengeDetailsPage
)

type Ui struct {
	common     common.Common
	pages      []common.Page
	activePage page
	footer     *footer.Footer
	showFooter bool
}

func NewUi(c common.Common) *Ui {
	ui := &Ui{
		common:     c,
		pages:      make([]common.Page, 2),
		activePage: challengesPage,
		showFooter: true,
	}

	ui.footer = footer.New(c, ui)
	return ui
}

func (ui *Ui) getMargins() (wm, hm int) {
	style := ui.common.Styles.App
	wm = style.GetHorizontalFrameSize()
	hm = style.GetVerticalFrameSize()
	if ui.showFooter {
		hm += ui.footer.Height()
	}
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
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, ui.common.KeyMap.Quit):
			return ui, tea.Quit
		case ui.activePage == challengeDetailsPage && key.Matches(msg, ui.common.KeyMap.Back):
			ui.activePage = challengesPage
		case key.Matches(msg, ui.common.KeyMap.HideFooter):
			cmds = append(cmds, footer.ToggleFooterCmd)
		case key.Matches(msg, ui.common.KeyMap.Help):
			cmds = append(cmds, footer.ToggleHelpCmd)
		}
	case selector.SelectMsg:
		switch msg.IdentifiableItem.(type) {
		case challenges.Item:
			challenge := msg.IdentifiableItem.(challenges.Item).Challenge
			cmds = append(cmds, ui.selectChallengeCmd(challenge))
		}
	case commands.SelectChallengeMsg:
		ui.activePage = challengeDetailsPage
	case footer.ToggleFooterMsg:
		ui.showFooter = !ui.showFooter
	case footer.ToggleHelpMsg:
		ui.footer.SetShowAll(!ui.footer.ShowAll())
	}
	f, cmd := ui.footer.Update(msg)
	ui.footer = f.(*footer.Footer)
	if cmd != nil {
		cmds = append(cmds, cmd)
	}

	m, cmd := ui.pages[ui.activePage].Update(msg)
	ui.pages[ui.activePage] = m.(common.Page)
	if cmd != nil {
		cmds = append(cmds, cmd)
	}

	ui.SetSize(ui.common.Width, ui.common.Height)
	return ui, tea.Batch(cmds...)
}

func (ui *Ui) View() string {
	view := ui.pages[ui.activePage].View()
	if ui.showFooter {
		view = lipgloss.JoinVertical(lipgloss.Left, view, ui.footer.View())
	}
	return ui.common.Styles.App.Render(view)
}

func (ui *Ui) selectChallengeCmd(challenge challenge.Challenge) tea.Cmd {
	return func() tea.Msg {
		return commands.SelectChallengeMsg(challenge)
	}
}

func (ui *Ui) ShortHelp() []key.Binding {
	bindings := make([]key.Binding, 0)
	bindings = append(bindings, ui.pages[ui.activePage].ShortHelp()...)
	bindings = append(bindings, ui.common.KeyMap.Help, ui.common.KeyMap.HideFooter, ui.common.KeyMap.Quit)

	return bindings
}

func (ui *Ui) FullHelp() [][]key.Binding {
	bindings := make([][]key.Binding, 0)

	bindings = append(bindings, ui.pages[ui.activePage].FullHelp()...)
	bindings = append(bindings,
		[]key.Binding{
			ui.common.KeyMap.Help,
			ui.common.KeyMap.HideFooter,
			ui.common.KeyMap.Quit,
		},
	)

	return bindings
}
