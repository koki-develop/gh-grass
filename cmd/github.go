package cmd

import (
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
	Weeks []struct {
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

func fetchCalendar(username string) (calendar, error) {
	client, err := gh.GQLClient(nil)
	if err != nil {
		return calendar{}, err
	}

	if flagUser == "" {
		var query viewerQuery
		if err := client.Query("contributions", &query, nil); err != nil {
			return calendar{}, err
		}
		return query.Viewer.ContributionsCollection.ContributionCalendar, nil
	} else {
		var query userQuery
		if err := client.Query("contributions", &query, map[string]interface{}{"user": graphql.String(username)}); err != nil {
			return calendar{}, err
		}
		return query.User.ContributionsCollection.ContributionCalendar, nil
	}
}
