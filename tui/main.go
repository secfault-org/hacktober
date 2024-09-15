package main

import (
	"fmt"
	"os"
	"time"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type Challenge struct {
	title       string
	description string
	completed   bool
	releaseDate time.Time
}

func (c Challenge) unlocked() bool {
	return time.Now().After(c.releaseDate)
}

func (c Challenge) Title() string {
	if c.unlocked() {
		if c.completed {
			return fmt.Sprintf("%s 🌟", c.title)
		} else {
			return c.title
		}
	} else {
		return lipgloss.NewStyle().Foreground(lipgloss.Color("241")).Render(c.title)
	}
}
func (c Challenge) Description() string {
	if c.unlocked() {
		return c.description
	} else {
		return "🔒"
	}
}
func (c Challenge) FilterValue() string { return c.title }

type model struct {
	list list.Model
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "q" {
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return docStyle.Render(m.list.View())
}

func main() {
	itmes := []list.Item{
		Challenge{title: "Challenge 1", description: "Description 1", completed: true},
		Challenge{title: "Challenge 2", description: "Description 2", completed: true},
		Challenge{title: "Challenge 3", description: "Description 3", releaseDate: time.Now().Add(-12 * time.Hour)},
		Challenge{title: "Challenge 4", description: "Description 4", releaseDate: time.Now().Add(24 * time.Hour)},
	}

	m := model{list: list.New(itmes, list.NewDefaultDelegate(), 0, 0)}
	m.list.Title = "Challenges"

	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
