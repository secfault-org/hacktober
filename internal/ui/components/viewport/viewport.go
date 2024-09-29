package viewport

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/secfault-org/hacktober/internal/ui/common"
)

type Viewport struct {
	common common.Common
	*viewport.Model
}

// New returns a new Viewport.
func New(c common.Common) *Viewport {
	vp := viewport.New(c.Width, c.Height)
	vp.MouseWheelEnabled = true
	return &Viewport{
		common: c,
		Model:  &vp,
	}
}

func (v *Viewport) SetSize(width, height int) {
	v.common.SetSize(width, height)
	v.Model.Width = width
	v.Model.Height = height
}

func (v *Viewport) Init() tea.Cmd {
	return nil
}

func (v *Viewport) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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

func (v *Viewport) HalfViewDown() {
	v.Model.HalfViewDown()
}

func (v *Viewport) HalfViewUp() {
	v.Model.HalfViewUp()
}

func (v *Viewport) ViewUp() []string {
	return v.Model.ViewUp()
}

func (v *Viewport) ViewDown() []string {
	return v.Model.ViewDown()
}

func (v *Viewport) LineUp(n int) []string {
	return v.Model.LineUp(n)
}

func (v *Viewport) LineDown(n int) []string {
	return v.Model.LineDown(n)
}

func (v *Viewport) ScrollPercent() float64 {
	return v.Model.ScrollPercent()
}
