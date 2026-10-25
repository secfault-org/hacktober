package viewport

import (
	"charm.land/bubbles/v2/key"
	"charm.land/bubbles/v2/viewport"
	tea "charm.land/bubbletea/v2"
	"github.com/secfault-org/hacktober/internal/ui/keymap"
)

type Viewport struct {
	*viewport.Model

	keyMap *keymap.KeyMap
}

// New returns a new Viewport.
func New(keyMap *keymap.KeyMap) *Viewport {
	vp := viewport.New()
	vp.MouseWheelEnabled = true
	return &Viewport{
		Model:  &vp,
		keyMap: keyMap,
	}
}

func (v *Viewport) SetSize(width, height int) {
	v.Model.SetWidth(width)
	v.Model.SetHeight(height)
}

func (v *Viewport) Init() tea.Cmd {
	return v.Model.Init()
}

func (v *Viewport) Update(msg tea.Msg) (*Viewport, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, v.keyMap.GotoTop):
			v.GotoTop()
		case key.Matches(msg, v.keyMap.GotoBottom):
			v.GotoBottom()
		}
	}
	vp, cmd := v.Model.Update(msg)
	v.Model = &vp
	return v, cmd
}

func (v *Viewport) View() string {
	return v.Model.View()
}

func (v *Viewport) SetContent(content string) {
	v.Model.SetContent(content)
}

func (v *Viewport) GotoTop() {
	v.Model.GotoTop()
}

func (v *Viewport) GotoBottom() {
	v.Model.GotoBottom()
}

func (v *Viewport) ScrollPercent() float64 {
	return v.Model.ScrollPercent()
}
