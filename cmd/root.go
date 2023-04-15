package cmd

import (
	"fmt"
	"os"

	"github.com/cli/go-gh"
	"github.com/spf13/cobra"
)

type viewerQuery struct {
	Viewer struct {
		ContributionsCollection struct {
			ContributionCalendar struct {
				Weeks []struct {
					ContributionDays []struct {
						Color string
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
