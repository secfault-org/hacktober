package viewport

import (
	"charm.land/bubbles/v2/key"
	"charm.land/bubbles/v2/viewport"
	tea "charm.land/bubbletea/v2"
	"github.com/secfault-org/hacktober/internal/ui/common"
)

type Viewport struct {
	common common.Common
	*viewport.Model
}

// New returns a new Viewport.
func New(c common.Common) *Viewport {
	vp := viewport.New(viewport.WithHeight(c.Height), viewport.WithWidth(c.Width))
	vp.MouseWheelEnabled = true
	return &Viewport{
		common: c,
		Model:  &vp,
	}
}

func (v *Viewport) SetSize(width, height int) {
	v.common.SetSize(width, height)
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
		case key.Matches(msg, v.common.KeyMap.GotoTop):
			v.GotoTop()
		case key.Matches(msg, v.common.KeyMap.GotoBottom):
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
