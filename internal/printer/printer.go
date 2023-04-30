package printer

import (
	"fmt"
	"io"

	"github.com/charmbracelet/lipgloss"
	"github.com/koki-develop/gh-grass/internal/github"
)

type PrintOptions struct {
	Theme    Theme
	Calendar github.Calendar
	Grass    string
}

func PrintGrass(w io.Writer, options PrintOptions) error {
	grasses := []string{}

	for _, week := range options.Calendar.Weeks {
		for _, d := range week.ContributionDays {
			c := lipgloss.Color(options.Theme[d.ContributionLevel])
			style := lipgloss.NewStyle().Foreground(c)
			grasses = append(grasses, style.Render(options.Grass))
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
