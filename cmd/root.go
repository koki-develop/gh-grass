package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var (
	flagUser  string
	flagTheme string
)

var rootCmd = &cobra.Command{
	Use:          "gh grass",
	Short:        "Print github grass to console",
	Long:         "Print github grass to console.",
	Args:         cobra.ExactArgs(0),
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		t, ok := themes[flagTheme]
		if !ok {
			return fmt.Errorf("valid themes: %s", listThemes())
		}

		cal, err := fetchCalendar(flagUser)
		if err != nil {
			return err
		}

		if err := printGrass(os.Stdout, t, cal); err != nil {
			return err
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
	rootCmd.Flags().StringVarP(&flagTheme, "theme", "t", "dark", fmt.Sprintf("grass theme (%s)", strings.Join(listThemes(), "|")))
}
