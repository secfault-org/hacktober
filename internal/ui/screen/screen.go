package screen

import "github.com/secfault-org/hacktober/internal/ui/styles"

type Screen struct {
	Width, Height int
	theme         *styles.Styles
}

func New(w, h int) *Screen { return &Screen{Width: w, Height: h, theme: styles.DefaultStyles()} }

func (s *Screen) Theme() *styles.Styles         { return s.theme }
func (s *Screen) SetTheme(theme *styles.Styles) { s.theme = theme }
