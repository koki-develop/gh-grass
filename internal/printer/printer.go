package printer

import (
	"fmt"
	"io"

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

	rows := 7
	columns := (len(grasses) + rows - 1) / rows

	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			index := j*rows + i
			if index < len(grasses) {
				if j != 0 {
					fmt.Fprint(w, " ")
				}
				fmt.Fprint(w, grasses[index])
			}
		}
		if i < rows-1 && i < len(grasses)-1 {
			fmt.Fprintln(w)
		}
	}

	fmt.Fprintln(w)

	return nil
}
