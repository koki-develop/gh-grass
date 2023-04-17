package cmd

import (
	"time"

	"github.com/cli/go-gh"
	graphql "github.com/cli/shurcooL-graphql"
)

type contributionLevel string

const (
	contributionLevelNone           contributionLevel = "NONE"
	contributionLevelFirstQuartile  contributionLevel = "FIRST_QUARTILE"
	contributionLevelSecondQuartile contributionLevel = "SECOND_QUARTILE"
	contributionLevelThirdQuartile  contributionLevel = "THIRD_QUARTILE"
	contributionLevelFourthQuartile contributionLevel = "FOURTH_QUARTILE"
)

type DateTime struct {
	time.Time
}

type calendar struct {
	TotalContributions int
	Weeks              []struct {
		ContributionDays []struct {
			ContributionLevel contributionLevel
		}
	}
}

type fetchCalendarParameters struct {
	User *string
	From *time.Time
	To   *time.Time
}

func fetchCalendar(params fetchCalendarParameters) (calendar, error) {
	client, err := gh.GQLClient(nil)
	if err != nil {
		return calendar{}, err
	}

	v := map[string]interface{}{}
	if params.User != nil {
		v["user"] = graphql.String(*params.User)
	}
	if params.From != nil {
		v["from"] = DateTime{*params.From}
	}
	if params.To != nil {
		v["to"] = DateTime{*params.To}
	}

	var cal calendar

	switch {
	case params.User != nil && params.From != nil && params.To != nil:
		var q struct {
			User struct {
				ContributionsCollection struct {
					ContributionCalendar calendar
				} `graphql:"contributionsCollection(from: $from, to: $to)"`
			} `graphql:"user(login: $user)"`
		}
		err = client.Query("contributions", &q, v)
		cal = q.User.ContributionsCollection.ContributionCalendar
	case params.User != nil && params.From != nil:
		var q struct {
			User struct {
				ContributionsCollection struct {
					ContributionCalendar calendar
				} `graphql:"contributionsCollection(from: $from)"`
			} `graphql:"user(login: $user)"`
		}
		err = client.Query("contributions", &q, v)
		cal = q.User.ContributionsCollection.ContributionCalendar
	case params.User != nil && params.To != nil:
		var q struct {
			User struct {
				ContributionsCollection struct {
					ContributionCalendar calendar
				} `graphql:"contributionsCollection(to: $to)"`
			} `graphql:"user(login: $user)"`
		}
		err = client.Query("contributions", &q, v)
		cal = q.User.ContributionsCollection.ContributionCalendar
	case params.User != nil:
		var q struct {
			User struct {
				ContributionsCollection struct {
					ContributionCalendar calendar
				}
			} `graphql:"user(login: $user)"`
		}
		err = client.Query("contributions", &q, v)
		cal = q.User.ContributionsCollection.ContributionCalendar
	case params.User == nil && params.From != nil && params.To != nil:
		var q struct {
			Viewer struct {
				ContributionsCollection struct {
					ContributionCalendar calendar
				} `graphql:"contributionsCollection(from: $from, to: $to)"`
			}
		}
		err = client.Query("calendar", &q, v)
		cal = q.Viewer.ContributionsCollection.ContributionCalendar
	case params.User == nil && params.From != nil:
		var q struct {
			Viewer struct {
				ContributionsCollection struct {
					ContributionCalendar calendar
				} `graphql:"contributionsCollection(from: $from)"`
			}
		}
		err = client.Query("calendar", &q, v)
		cal = q.Viewer.ContributionsCollection.ContributionCalendar
	case params.User == nil && params.To != nil:
		var q struct {
			Viewer struct {
				ContributionsCollection struct {
					ContributionCalendar calendar
				} `graphql:"contributionsCollection(to: $to)"`
			}
		}
		err = client.Query("calendar", &q, v)
		cal = q.Viewer.ContributionsCollection.ContributionCalendar
	case params.User == nil:
		var q struct {
			Viewer struct {
				ContributionsCollection struct {
					ContributionCalendar calendar
				}
			}
		}
		err = client.Query("calendar", &q, v)
		cal = q.Viewer.ContributionsCollection.ContributionCalendar
	}

	return cal, err
}
