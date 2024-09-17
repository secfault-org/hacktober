package ui

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"hacktopber2024/pkg/model"
	"hacktopber2024/pkg/ui/common"
	"hacktopber2024/pkg/ui/components/selector"
	"hacktopber2024/pkg/ui/pages/challenge_detail"
	"hacktopber2024/pkg/ui/pages/challenges"
	"time"
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
}

func NewUi(c common.Common) *Ui {
	return &Ui{
		common:     c,
		pages:      make([]common.Page, 2),
		activePage: challengesPage,
	}
}

func (ui *Ui) getMargins() (wm, hm int) {
	style := ui.common.Styles.App
	wm = style.GetHorizontalFrameSize()
	hm = style.GetVerticalFrameSize()
	return
}

func (ui *Ui) SetSize(width, height int) {
	ui.common.SetSize(width, height)
	wm, hm := ui.getMargins()
	for _, p := range ui.pages {
		if p != nil {
			p.SetSize(width-wm, height-hm)
		}
	}
}

func (ui *Ui) Init() tea.Cmd {
	items := []selector.IdentifiableItem{
		challenges.Item{Challenge: model.Challenge{Name: "Challenge 1", Description: "Description 1", ReleaseDate: time.Now().Add(-12 * time.Hour)}},
		challenges.Item{Challenge: model.Challenge{Name: "Challenge 2", Description: "Description 2", ReleaseDate: time.Now().Add(12 * time.Hour)}},
		challenges.Item{Challenge: model.Challenge{Name: "Challenge 3", Description: "Description 3", ReleaseDate: time.Now().Add(48 * time.Hour)}},
		challenges.Item{Challenge: model.Challenge{Name: "Challenge 4", Description: "Description 4", ReleaseDate: time.Now().Add(96 * time.Hour)}},
	}
	ui.pages[challengesPage] = challenges.New(ui.common, items)
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
		}
	case selector.SelectMsg:
		switch msg.IdentifiableItem.(type) {
		case challenges.Item:
			challenge := msg.IdentifiableItem.(challenges.Item).Challenge
			cmds = append(cmds, ui.selectChallengeCmd(challenge))
		}
	case challenge_detail.SelectChallengeMsg:
		ui.activePage = challengeDetailsPage
	}
	m, cmd := ui.pages[ui.activePage].Update(msg)
	ui.pages[ui.activePage] = m.(common.Page)
	if cmd != nil {
		cmds = append(cmds, cmd)
	}

	return ui, tea.Batch(cmds...)
}

func (ui *Ui) View() string {
	return ui.common.Styles.App.Render(ui.pages[ui.activePage].View())
}

func (ui *Ui) selectChallengeCmd(challenge model.Challenge) tea.Cmd {
	return func() tea.Msg {
		return challenge_detail.SelectChallengeMsg(challenge)
	}
}
