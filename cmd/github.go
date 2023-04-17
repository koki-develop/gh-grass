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

type contributions struct {
	ContributionsCollection struct {
		ContributionCalendar calendar
	}
}

type calendar struct {
	TotalContributions int
	Weeks              []struct {
		ContributionDays []struct {
			ContributionLevel contributionLevel
		}
	}
}

type viewerQuery struct {
	Viewer contributions
}

type userQuery struct {
	User contributions `graphql:"user(login: $user)"`
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

	if params.User == nil {
		var query viewerQuery
		if err := client.Query("contributions", &query, v); err != nil {
			return calendar{}, err
		}
		return query.Viewer.ContributionsCollection.ContributionCalendar, nil
	} else {
		var query userQuery
		if err := client.Query("contributions", &query, v); err != nil {
			return calendar{}, err
		}
		return query.User.ContributionsCollection.ContributionCalendar, nil
	}
}
