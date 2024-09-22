package main

import (
	"context"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/secfault-org/hacktober/pkg/repository"
	"github.com/secfault-org/hacktober/pkg/ui"
	"github.com/secfault-org/hacktober/pkg/ui/common"
	"os"
)

func main() {
	ctx := context.Background()
	ctx = repository.WithContext(ctx, repository.NewRepository(ctx, "challenges/2024"))
	c := common.NewCommon(ctx, nil, 80, 24)

	app := ui.NewUi(c)

	p := tea.NewProgram(app, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
