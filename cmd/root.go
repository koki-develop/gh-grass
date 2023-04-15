package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/cli/go-gh"
	"github.com/spf13/cobra"
)

type contributionLevel string

const (
	contributionLevelNone           contributionLevel = "NONE"
	contributionLevelFirstQuartile  contributionLevel = "FIRST_QUARTILE"
	contributionLevelSecondQuartile contributionLevel = "SECOND_QUARTILE"
	contributionLevelThirdQuartile  contributionLevel = "THIRD_QUARTILE"
	contributionLevelFourthQuartile contributionLevel = "FOURTH_QUARTILE"
)

var themes = map[string]map[contributionLevel]string{
	"dark-default": {
		contributionLevelNone:           "#161B22",
		contributionLevelFirstQuartile:  "#0E4429",
		contributionLevelSecondQuartile: "#006D32",
		contributionLevelThirdQuartile:  "#26A641",
		contributionLevelFourthQuartile: "#39D353",
	},
}

type viewerQuery struct {
	Viewer struct {
		ContributionsCollection struct {
			ContributionCalendar struct {
				Weeks []struct {
					ContributionDays []struct {
						ContributionLevel contributionLevel
					}
				}
			}
		}
	}
}

var rootCmd = &cobra.Command{
	Use: "grass",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := gh.GQLClient(nil)
		if err != nil {
			return err
		}

		var query viewerQuery
		if err := client.Query("contributions", &query, nil); err != nil {
			return err
		}

		theme := themes["dark-default"] // TODO: from flag

		for i := 0; i < 7; i++ {
			chars := make([]string, len(query.Viewer.ContributionsCollection.ContributionCalendar.Weeks))
			for j, w := range query.Viewer.ContributionsCollection.ContributionCalendar.Weeks {
				d := w.ContributionDays[i]
				c := lipgloss.Color(theme[d.ContributionLevel])
				style := lipgloss.NewStyle().Foreground(c)
				chars[j] = style.Render("■")
			}
			fmt.Print(strings.Join(chars, " "))
			fmt.Print("\n")
		}
		return nil
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
