package cmd

import (
	"fmt"
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

var (
	flagUser  string
	flagTheme string
)

var rootCmd = &cobra.Command{
	Use:   "gh grass",
	Short: "Print github grass to console",
	Long:  "Print github grass to console.",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		t, ok := themes[flagTheme]
		if !ok {
			valid := []string{}
			for k := range themes {
				valid = append(valid, k)
			}
			return fmt.Errorf("valid themes: %s", valid)
		}

		cal, err := fetchCalendar(flagUser)
		if err != nil {
			return err
		}

		for i := 0; i < 7; i++ {
			for j, w := range cal.Weeks {
				d := w.ContributionDays[i]
				c := lipgloss.Color(t[d.ContributionLevel])
				style := lipgloss.NewStyle().Foreground(c)
				fmt.Print(style.Render("■"))

				if j+1 != len(cal.Weeks) {
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
	rootCmd.Flags().StringVarP(&flagTheme, "theme", "t", "dark", "grass theme (dark|light)")
}
