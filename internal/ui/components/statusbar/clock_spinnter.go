package statusbar

import (
	"time"

	"charm.land/bubbles/v2/spinner"
)

var (
	Clock = spinner.Spinner{
		Frames: []string{"🕐", "🕑", "🕒", "🕓", "🕔", "🕕", "🕖", "🕗", "🕘", "🕙", "🕚", "🕛"},
		FPS:    time.Second,
	}
)
