package confetti

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/harmonica"
	"github.com/charmbracelet/lipgloss"
	"github.com/secfault-org/hacktober/internal/ui/common"
	"math/rand"
	"time"
)

func Sample[T any](arr []T) T {
	return arr[rand.Intn(len(arr))]
}

const (
	framesPerSecond = 30.0
	numParticles    = 75
)

var (
	colors     = []string{"#a864fd", "#29cdff", "#78ff44", "#ff718d", "#fdff6a"}
	characters = []string{"█", "▓", "▒", "░", "▄", "▀"}
)

type frameMsg time.Time

func animate() tea.Cmd {
	return tea.Tick(time.Second/framesPerSecond, func(t time.Time) tea.Msg {
		return frameMsg(t)
	})
}

type Model struct {
	common common.Common
	system *System
}

func Spawn(width, height int) []*Particle {
	particles := []*Particle{}
	for i := 0; i < numParticles; i++ {
		x := float64(width / 2)
		y := float64(height)

		p := Particle{
			Physics: harmonica.NewProjectile(
				harmonica.FPS(framesPerSecond),
				harmonica.Point{X: x + (float64(width/4) * (rand.Float64() - 0.5)), Y: y, Z: 0},
				harmonica.Vector{X: (rand.Float64() - 0.5) * 200, Y: rand.Float64() * -200, Z: 0},
				harmonica.Vector{Y: 30},
			),
			Char: lipgloss.NewStyle().
				Foreground(lipgloss.Color(Sample(colors))).
				Render(Sample(characters)),
		}

		particles = append(particles, &p)
	}
	return particles
}

func InitialModel(cmn common.Common) Model {
	return Model{
		common: cmn,
		system: &System{
			Particles: []*Particle{},
			Frame:     Frame{},
		},
	}
}

// Init initializes the confetti after a small delay
func (m Model) Init() tea.Cmd {
	m.system.Particles = Spawn(m.common.Width, m.common.Height)
	return animate()
}

// Update updates the model every frame, it handles the animation loop and
// updates the particle physics every frame
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case frameMsg:
		m.system.Update()
		return m, animate()
	case tea.WindowSizeMsg:
		m.system.Frame.Width = msg.Width
		m.system.Frame.Height = msg.Height
		return m, nil
	default:
		return m, nil
	}
}

// View displays all the particles on the screen
func (m Model) View() string {
	return m.system.Render()
}

func (m Model) Reset() {
	m.system.Particles = Spawn(m.common.Width, m.common.Height)
}
