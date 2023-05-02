package printer

import "github.com/koki-develop/gh-grass/internal/github"

type Theme map[github.ContributionLevel]string

var Themes = map[string]Theme{
	"dark": {
		github.ContributionLevelNone:           "#161B22",
		github.ContributionLevelFirstQuartile:  "#0E4429",
		github.ContributionLevelSecondQuartile: "#006D32",
		github.ContributionLevelThirdQuartile:  "#26A641",
		github.ContributionLevelFourthQuartile: "#39D353",
	},
	"light": {
		github.ContributionLevelNone:           "#EBEDF0",
		github.ContributionLevelFirstQuartile:  "#9BE9A8",
		github.ContributionLevelSecondQuartile: "#40C463",
		github.ContributionLevelThirdQuartile:  "#30A14E",
		github.ContributionLevelFourthQuartile: "#216E39",
	},
}

func GetTheme(name string) (Theme, bool) {
	t, ok := Themes[name]
	return t, ok
}

func ListThemes() []string {
	ts := []string{}

	for t := range Themes {
		ts = append(ts, t)
	}

	return ts
}
