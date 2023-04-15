package cmd

type theme map[contributionLevel]string

var themes = map[string]theme{
	"dark": {
		contributionLevelNone:           "#161B22",
		contributionLevelFirstQuartile:  "#0E4429",
		contributionLevelSecondQuartile: "#006D32",
		contributionLevelThirdQuartile:  "#26A641",
		contributionLevelFourthQuartile: "#39D353",
	},
	"light": {
		contributionLevelNone:           "#EBEDF0",
		contributionLevelFirstQuartile:  "#9BE9A8",
		contributionLevelSecondQuartile: "#40C463",
		contributionLevelThirdQuartile:  "#30A14E",
		contributionLevelFourthQuartile: "#216E39",
	},
}

func listThemes() []string {
	ts := []string{}

	for t := range themes {
		ts = append(ts, t)
	}

	return ts
}
