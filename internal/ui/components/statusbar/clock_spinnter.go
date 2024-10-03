package statusbar

import (
	"github.com/charmbracelet/bubbles/spinner"
	"time"
)

var (
	Clock = spinner.Spinner{
		Frames: []string{"🕐", "🕑", "🕒", "🕓", "🕔", "🕕", "🕖", "🕗", "🕘", "🕙", "🕚", "🕛"},
		FPS:    time.Second,
	}
)
