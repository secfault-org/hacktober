package confetti

import (
	"fmt"
	"github.com/charmbracelet/harmonica"
	"github.com/charmbracelet/lipgloss"
	"strings"
)

type System struct {
	Frame     Frame
	Particles []*Particle
}

type Particle struct {
	Char    string
	Color   lipgloss.Color
	Physics *harmonica.Projectile
	Hidden  bool
}

type Frame struct {
	Width  int
	Height int
}

func RemoveParticleFromArray(s []*Particle, i int) []*Particle {
	s[i] = nil
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func (s *System) Update() {
	for i := len(s.Particles) - 1; i >= 0; i-- {
		p := s.Particles[i]
		pos := p.Physics.Position()

		// remove particles that are hidden or out of the side/bottom of the frame
		if p.Hidden || pos.X > float64(s.Frame.Width) || pos.X < 0 || pos.Y > float64(s.Frame.Height) {
			s.Particles = RemoveParticleFromArray(s.Particles, i)
		} else {
			s.Particles[i].Physics.Update()
		}
	}
}

func (s *System) Visible(p *Particle) bool {
	pos := p.Physics.Position()
	x := int(pos.X)
	y := int(pos.Y)
	return !p.Hidden && y >= 0 && y < s.Frame.Height-1 && x >= 0 && x < s.Frame.Width-1
}

func (s *System) Render() string {
	var out strings.Builder
	plane := make([][]string, s.Frame.Height)
	for i := range plane {
		plane[i] = make([]string, s.Frame.Width)
	}
	for _, p := range s.Particles {
		if s.Visible(p) {
			pos := p.Physics.Position()
			plane[int(pos.Y)][int(pos.X)] = p.Char
		}
	}
	for i := range plane {
		for _, col := range plane[i] {
			if col == "" {
				fmt.Fprint(&out, " ")
			} else {
				fmt.Fprint(&out, col)
			}
		}
		fmt.Fprint(&out, "\n")
	}
	return out.String()
}
