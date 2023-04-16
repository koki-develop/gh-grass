package cmd

import (
	"fmt"
	"io"

	"github.com/charmbracelet/lipgloss"
)

func printGrass(w io.Writer, t theme, cal calendar) error {
	for i := 0; i < 7; i++ {
		for j, week := range cal.Weeks {
			if len(week.ContributionDays) < i+1 {
				continue
			}

			d := week.ContributionDays[i]
			c := lipgloss.Color(t[d.ContributionLevel])
			style := lipgloss.NewStyle().Foreground(c)

			if _, err := fmt.Fprint(w, style.Render("■")); err != nil {
				return err
			}

			if j+1 != len(cal.Weeks) {
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
