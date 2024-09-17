package main

import (
	"context"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"hacktopber2024/pkg/ui"
	"hacktopber2024/pkg/ui/common"
	"os"
)

func main() {
	c := common.NewCommon(context.Background(), nil, 80, 24)

	app := ui.NewUi(c)

	p := tea.NewProgram(app, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
