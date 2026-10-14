package common

import tea "charm.land/bubbletea/v2"

// ErrorMsg is a Bubble Tea message that represents an error.
type ErrorMsg error

// ErrorCmd returns an ErrorMsg from error.
func ErrorCmd(err error) tea.Cmd {
	return func() tea.Msg {
		return ErrorMsg(err)
	}
}
