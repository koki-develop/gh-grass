package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

var (
	flagUser  string
	flagTheme string
	flagGrass string
	flagTotal bool
)

var rootCmd = &cobra.Command{
	Use:          "gh grass",
	Short:        "Grow github grass to console",
	Long:         "Grow github grass to console.",
	Args:         cobra.ExactArgs(0),
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		t, ok := themes[flagTheme]
		if !ok {
			return fmt.Errorf("valid themes: %s", listThemes())
		}

		params := fetchCalendarParameters{}
		if flagUser != "" {
			params.User = &flagUser
		}

		cal, err := fetchCalendar(params)
		if err != nil {
			return err
		}

		if flagTotal {
			p := message.NewPrinter(language.English)
			p.Printf("%d contributions in the last year\n", cal.TotalContributions)
		}

		if err := printGrass(os.Stdout, printOptions{theme: t, calendar: cal, grass: flagGrass}); err != nil {
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
	rootCmd.Flags().StringVarP(&flagGrass, "grass", "g", "■", "grass string")
	rootCmd.Flags().BoolVar(&flagTotal, "total", false, "print total contributions")
}
