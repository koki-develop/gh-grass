package cmd

import (
	"fmt"
	"io"

	"github.com/charmbracelet/lipgloss"
	"github.com/koki-develop/gh-grass/internal/github"
)

type printOptions struct {
	theme    theme
	calendar github.Calendar
	grass    string
}

func printGrass(w io.Writer, options printOptions) error {
	grasses := []string{}

	for _, week := range options.calendar.Weeks {
		for _, d := range week.ContributionDays {
			c := lipgloss.Color(options.theme[d.ContributionLevel])
			style := lipgloss.NewStyle().Foreground(c)
			grasses = append(grasses, style.Render(options.grass))
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
