package cmd

import (
	"fmt"
	"os"

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

		fmt.Printf("%#v", query)
		return nil
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
