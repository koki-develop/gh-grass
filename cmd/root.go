package cmd

import (
	"fmt"
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/cli/go-gh"
	graphql "github.com/cli/shurcooL-graphql"
	"github.com/spf13/cobra"
)

var (
	flagUser  string
	flagTheme string
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

type contributions struct {
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

type viewerQuery struct {
	Viewer contributions
}

type userQuery struct {
	User contributions `graphql:"user(login: $user)"`
}

var rootCmd = &cobra.Command{
	Use:   "gh grass",
	Short: "Print github grass to console",
	Long:  "Print github grass to console.",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := gh.GQLClient(nil)
		if err != nil {
			return err
		}

		var contributions contributions
		if flagUser == "" {
			var query viewerQuery
			if err := client.Query("contributions", &query, nil); err != nil {
				return err
			}
			contributions = query.Viewer
		} else {
			var query userQuery
			if err := client.Query("contributions", &query, map[string]interface{}{"user": graphql.String(flagUser)}); err != nil {
				return err
			}
			contributions = query.User
		}

		theme, ok := themes[flagTheme]
		if !ok {
			valid := []string{}
			for k := range themes {
				valid = append(valid, k)
			}
			return fmt.Errorf("valid themes: %s", valid)
		}

		for i := 0; i < 7; i++ {
			for j, w := range contributions.ContributionsCollection.ContributionCalendar.Weeks {
				d := w.ContributionDays[i]
				c := lipgloss.Color(theme[d.ContributionLevel])
				style := lipgloss.NewStyle().Foreground(c)
				fmt.Print(style.Render("■"))

				if j+1 != len(contributions.ContributionsCollection.ContributionCalendar.Weeks) {
					fmt.Print(" ")
				}
			}
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

func init() {
	rootCmd.Flags().SortFlags = false

	rootCmd.Flags().StringVarP(&flagUser, "user", "u", "", "github username")
	rootCmd.Flags().StringVarP(&flagTheme, "theme", "t", "dark-default", "grass theme")
}
