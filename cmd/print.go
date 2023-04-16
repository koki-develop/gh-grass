package cmd

import (
	"fmt"
	"io"

	"github.com/charmbracelet/lipgloss"
)

type printOptions struct {
	theme    theme
	calendar calendar
	grass    string
}

func printGrass(w io.Writer, options printOptions) error {
	for i := 0; i < 7; i++ {
		for j, week := range options.calendar.Weeks {
			if len(week.ContributionDays) < i+1 {
				continue
			}

			d := week.ContributionDays[i]
			c := lipgloss.Color(options.theme[d.ContributionLevel])
			style := lipgloss.NewStyle().Foreground(c)

			if _, err := fmt.Fprint(w, style.Render(options.grass)); err != nil {
				return err
			}

			if j+1 != len(options.calendar.Weeks) {
				if _, err := fmt.Fprint(w, " "); err != nil {
					return err
				}
			}
		}
		if _, err := fmt.Fprint(w, "\n"); err != nil {
			return err
		}
	}

	return nil
}
