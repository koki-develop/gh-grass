package github

import (
	"time"

	"github.com/cli/go-gh/v2/pkg/api"
	graphql "github.com/cli/shurcooL-graphql"
)

type ContributionLevel string

const (
	ContributionLevelNone           ContributionLevel = "NONE"
	ContributionLevelFirstQuartile  ContributionLevel = "FIRST_QUARTILE"
	ContributionLevelSecondQuartile ContributionLevel = "SECOND_QUARTILE"
	ContributionLevelThirdQuartile  ContributionLevel = "THIRD_QUARTILE"
	ContributionLevelFourthQuartile ContributionLevel = "FOURTH_QUARTILE"
)

type DateTime struct {
	time.Time
}

type Calendar struct {
	TotalContributions int
	Weeks              []struct {
		ContributionDays []struct {
			ContributionLevel ContributionLevel
		}
	}
}

type FetchCalendarParameters struct {
	User *string
	From *time.Time
	To   *time.Time
}

func FetchCalendar(params FetchCalendarParameters) (Calendar, error) {
	client, err := api.DefaultGraphQLClient()
	if err != nil {
		return Calendar{}, err
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

	var cal Calendar

	switch {
	case params.User != nil && params.From != nil && params.To != nil:
		var q struct {
			User struct {
				ContributionsCollection struct {
					ContributionCalendar Calendar
				} `graphql:"contributionsCollection(from: $from, to: $to)"`
			} `graphql:"user(login: $user)"`
		}
		err = client.Query("contributions", &q, v)
		cal = q.User.ContributionsCollection.ContributionCalendar
	case params.User != nil && params.From != nil:
		var q struct {
			User struct {
				ContributionsCollection struct {
					ContributionCalendar Calendar
				} `graphql:"contributionsCollection(from: $from)"`
			} `graphql:"user(login: $user)"`
		}
		err = client.Query("contributions", &q, v)
		cal = q.User.ContributionsCollection.ContributionCalendar
	case params.User != nil && params.To != nil:
		var q struct {
			User struct {
				ContributionsCollection struct {
					ContributionCalendar Calendar
				} `graphql:"contributionsCollection(to: $to)"`
			} `graphql:"user(login: $user)"`
		}
		err = client.Query("contributions", &q, v)
		cal = q.User.ContributionsCollection.ContributionCalendar
	case params.User != nil:
		var q struct {
			User struct {
				ContributionsCollection struct {
					ContributionCalendar Calendar
				}
			} `graphql:"user(login: $user)"`
		}
		err = client.Query("contributions", &q, v)
		cal = q.User.ContributionsCollection.ContributionCalendar
	case params.User == nil && params.From != nil && params.To != nil:
		var q struct {
			Viewer struct {
				ContributionsCollection struct {
					ContributionCalendar Calendar
				} `graphql:"contributionsCollection(from: $from, to: $to)"`
			}
		}
		err = client.Query("calendar", &q, v)
		cal = q.Viewer.ContributionsCollection.ContributionCalendar
	case params.User == nil && params.From != nil:
		var q struct {
			Viewer struct {
				ContributionsCollection struct {
					ContributionCalendar Calendar
				} `graphql:"contributionsCollection(from: $from)"`
			}
		}
		err = client.Query("calendar", &q, v)
		cal = q.Viewer.ContributionsCollection.ContributionCalendar
	case params.User == nil && params.To != nil:
		var q struct {
			Viewer struct {
				ContributionsCollection struct {
					ContributionCalendar Calendar
				} `graphql:"contributionsCollection(to: $to)"`
			}
		}
		err = client.Query("calendar", &q, v)
		cal = q.Viewer.ContributionsCollection.ContributionCalendar
	case params.User == nil:
		var q struct {
			Viewer struct {
				ContributionsCollection struct {
					ContributionCalendar Calendar
				}
			}
		}
		err = client.Query("calendar", &q, v)
		cal = q.Viewer.ContributionsCollection.ContributionCalendar
	}

	return cal, err
}
