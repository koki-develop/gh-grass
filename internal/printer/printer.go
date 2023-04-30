package printer

import (
	"fmt"
	"io"

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
			style := lipgloss.NewStyle().Foreground(c)
			grasses = append(grasses, style.Render(p.grass))
		}
	}

	if p.animate {
		m := newModel(grasses)
		p := tea.NewProgram(m, tea.WithOutput(w))
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
	rows := 7
	columns := (len(grasses) + rows - 1) / rows

	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			index := j*rows + i
			if index < len(grasses) {
				if j != 0 {
					if _, err := fmt.Fprint(w, " "); err != nil {
						return err
					}
				}
				if _, err := fmt.Fprint(w, grasses[index]); err != nil {
					return err
				}
			}
		}
		if i < rows-1 && i < len(grasses)-1 {
			if _, err := fmt.Fprintln(w); err != nil {
				return err
			}
		}
	}

	if _, err := fmt.Fprintln(w); err != nil {
		return err
	}

	return nil
}
