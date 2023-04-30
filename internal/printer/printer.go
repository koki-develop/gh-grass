package printer

import (
	"fmt"
	"io"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/koki-develop/gh-grass/internal/github"
)

type Printer struct {
	theme   Theme
	grass   string
	animate bool
}

type Config struct {
	Theme   Theme
	Grass   string
	Animate bool
}

func New(cfg *Config) *Printer {
	return &Printer{
		theme:   cfg.Theme,
		grass:   cfg.Grass,
		animate: cfg.Animate,
	}
}

type PrintOptions struct {
	Theme    Theme
	Calendar github.Calendar
	Grass    string
}

func (p *Printer) Print(w io.Writer, calendar github.Calendar) error {
	grasses := []string{}

	for _, week := range calendar.Weeks {
		for _, d := range week.ContributionDays {
			c := lipgloss.Color(p.theme[d.ContributionLevel])
			style := lipgloss.NewStyle().Foreground(c).Bold(true)
			grasses = append(grasses, style.Render(p.grass))
		}
	}

	if p.animate {
		m := newModel(p, grasses)
		p := tea.NewProgram(m, tea.WithOutput(os.Stderr))
		if _, err := p.Run(); err != nil {
			return err
		}
	} else {
		if err := p.print(w, grasses); err != nil {
			return err
		}
	}

	return nil
}

func (p *Printer) print(w io.Writer, grasses []string) error {
	if _, err := fmt.Fprintln(w, p.grasses(grasses, len(grasses))); err != nil {
		return err
	}
	return nil
}

func (p *Printer) grasses(grasses []string, to int) string {
	var v strings.Builder

	rows := 7

	columns := (to + rows - 1) / rows
	for r := 0; r < rows; r++ {
		for c := 0; c < columns; c++ {
			pos := c*rows + r
			if pos < to {
				if c > 0 {
					v.WriteRune(' ')
				}
				v.WriteString(grasses[pos])
			}
		}
		v.WriteRune('\n')
	}

	return v.String()
}
