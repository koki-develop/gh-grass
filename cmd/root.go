package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/koki-develop/gh-grass/internal/github"
	"github.com/koki-develop/gh-grass/internal/printer"
	"github.com/koki-develop/gh-grass/internal/util"

	"github.com/jinzhu/now"
	"github.com/spf13/cobra"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

var (
	flagUser  string
	flagFrom  string
	flagTo    string
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
		themeName := util.GetEnvOr("GH_GRASS_THEME", flagTheme)
		if cmd.Flag("theme").Changed {
			themeName = flagTheme
		}

		t, ok := printer.Themes[themeName]
		if !ok {
			return fmt.Errorf("valid themes: %s", printer.ListThemes())
		}

		params := github.FetchCalendarParameters{}
		// --user
		if flagUser != "" {
			params.User = &flagUser
		}
		// --from
		if flagFrom != "" {
			t, err := now.Parse(flagFrom)
			if err != nil {
				return err
			}
			params.From = &t
		}
		// --to
		if flagTo != "" {
			f, err := now.Parse(flagTo)
			if err != nil {
				return err
			}
			params.To = &f
		}

		cal, err := github.FetchCalendar(params)
		if err != nil {
			return err
		}

		if flagTotal {
			p := message.NewPrinter(language.English)
			p.Printf("%d contributions in the last year\n", cal.TotalContributions)
		}

		p := printer.New(&printer.Config{
			Theme: t,
			Grass: flagGrass,
		})
		if err := p.Print(os.Stdout, cal); err != nil {
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
	rootCmd.Flags().StringVar(&flagFrom, "from", "", "only contributions made at this time or later will be counted")
	rootCmd.Flags().StringVar(&flagTo, "to", "", "only contributions made before and up to (including) this time will be counted")

	rootCmd.Flags().StringVarP(&flagTheme, "theme", "t", "dark", fmt.Sprintf("grass theme (%s)", strings.Join(printer.ListThemes(), "|")))
	rootCmd.Flags().StringVarP(&flagGrass, "grass", "g", "■", "grass string")
	rootCmd.Flags().BoolVar(&flagTotal, "total", false, "print total contributions")
}
