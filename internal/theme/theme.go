package theme

import "github.com/koki-develop/gh-grass/internal/github"

var (
	_ = Register("dark", "#161B22", "#0E4429", "#006D32", "#26A641", "#39D353")
	_ = Register("light", "#EBEDF0", "#9BE9A8", "#40C463", "#30A14E", "#216E39")
)

type Theme map[github.ContributionLevel]string

var registry = map[string]Theme{}

func Register(name, none, first, second, third, fourth string) Theme {
	t := Theme{
		github.ContributionLevelNone:           none,
		github.ContributionLevelFirstQuartile:  first,
		github.ContributionLevelSecondQuartile: second,
		github.ContributionLevelThirdQuartile:  third,
		github.ContributionLevelFourthQuartile: fourth,
	}
	registry[name] = t

	return t
}

func Get(name string) (Theme, bool) {
	t, ok := registry[name]
	return t, ok
}

func List() []string {
	ts := []string{}

	for t := range registry {
		ts = append(ts, t)
	}

	return ts
}
