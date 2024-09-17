package common

import (
	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
)

type Page interface {
	tea.Model
	help.KeyMap
	SetSize(width, height int)
}
