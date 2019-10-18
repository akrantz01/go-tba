package tba

import (
	"encoding/json"
	"github.com/akrantz01/go-tba/responses"
	"net/http"
	"strconv"
)

// Get a team's information by their key
func Team(team, apiKey string, opts *RequestOptions) (*responses.Team, int, error) {
	// Generate the request
	req := newRequest("/team/"+team, apiKey, opts)

	// Send the request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, 0, err
	}

	// Check the response type
	if resp.StatusCode == http.StatusUnauthorized {
		return nil, resp.StatusCode, ErrInvalidToken
	} else if resp.StatusCode == http.StatusNotModified {
		return nil, resp.StatusCode, nil
	} else if resp.StatusCode != 200 {
		return nil, resp.StatusCode, err
	}

	// Decode body
	var t responses.Team
	if err := json.NewDecoder(resp.Body).Decode(&t); err != nil {
		return nil, resp.StatusCode, err
	}

	return &t, resp.StatusCode, nil
}

// Get the simplified data of a team
func TeamSimple(team, apiKey string, opts *RequestOptions) (*responses.TeamSimple, int, error) {
	// Generate the request
	req := newRequest("/team/"+team+"/simple", apiKey, opts)

	// Send the request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, 0, err
	}

	// Check the response type
	if resp.StatusCode == http.StatusUnauthorized {
		return nil, resp.StatusCode, ErrInvalidToken
	} else if resp.StatusCode == http.StatusNotModified {
		return nil, resp.StatusCode, nil
	} else if resp.StatusCode != 200 {
		return nil, resp.StatusCode, err
	}

	// Decode body
	var t responses.TeamSimple
	if err := json.NewDecoder(resp.Body).Decode(&t); err != nil {
		return nil, resp.StatusCode, err
	}

	return &t, resp.StatusCode, nil
}

// Get the years in which a team participated in competition
func YearsTeamParticipated(team, apiKey string, opts *RequestOptions) ([]int, int, error) {
	// Generate the request
	req := newRequest("/team/"+team+"/years_participated", apiKey, opts)

	// Send the request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, 0, err
	}

	// Check the response type
	if resp.StatusCode == http.StatusUnauthorized {
		return nil, resp.StatusCode, ErrInvalidToken
	} else if resp.StatusCode == http.StatusNotModified {
		return nil, resp.StatusCode, nil
	} else if resp.StatusCode != 200 {
		return nil, resp.StatusCode, err
	}

	// Decode body
	var years []int
	if err := json.NewDecoder(resp.Body).Decode(&years); err != nil {
		return nil, resp.StatusCode, err
	}

	return years, resp.StatusCode, nil
}

// Get a list of the districts a team is in
func TeamDistricts(team, apiKey string, opts *RequestOptions) ([]responses.DistrictList, int, error) {
	// Generate the request
	req := newRequest("/team/"+team+"/districts", apiKey, opts)

	// Send the request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, 0, err
	}

	// Check the response type
	if resp.StatusCode == http.StatusUnauthorized {
		return nil, resp.StatusCode, ErrInvalidToken
	} else if resp.StatusCode == http.StatusNotModified {
		return nil, resp.StatusCode, nil
	} else if resp.StatusCode != 200 {
		return nil, resp.StatusCode, err
	}

	// Decode body
	var districts []responses.DistrictList
	if err := json.NewDecoder(resp.Body).Decode(&districts); err != nil {
		return nil, resp.StatusCode, err
	}

	return districts, resp.StatusCode, nil
}

// Get a list of named robots associated with a given team
func TeamRobots(team, apiKey string, opts *RequestOptions) ([]responses.TeamRobot, int, error) {
	// Generate the request
	req := newRequest("/team/"+team+"/robots", apiKey, opts)

	// Send the request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, 0, err
	}

	// Check the response type
	if resp.StatusCode == http.StatusUnauthorized {
		return nil, resp.StatusCode, ErrInvalidToken
	} else if resp.StatusCode == http.StatusNotModified {
		return nil, resp.StatusCode, nil
	} else if resp.StatusCode != 200 {
		return nil, resp.StatusCode, err
	}

	// Decode body
	var robots []responses.TeamRobot
	if err := json.NewDecoder(resp.Body).Decode(&robots); err != nil {
		return nil, resp.StatusCode, err
	}

	return robots, resp.StatusCode, nil
}

// Get a list of events a team is in
func TeamEvents(team, apiKey string, opts *RequestOptions) ([]responses.Event, int, error) {
	// Generate the request
	req := newRequest("/team/"+team+"/events", apiKey, opts)

	// Send the request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, 0, err
	}

	// Check the response type
	if resp.StatusCode == http.StatusUnauthorized {
		return nil, resp.StatusCode, ErrInvalidToken
	} else if resp.StatusCode == http.StatusNotModified {
		return nil, resp.StatusCode, nil
	} else if resp.StatusCode != 200 {
		return nil, resp.StatusCode, err
	}

	// Decode body
	var events []responses.Event
	if err := json.NewDecoder(resp.Body).Decode(&events); err != nil {
		return nil, resp.StatusCode, err
	}

	return events, resp.StatusCode, nil
}

// Get a simplified list of events a team is in
func TeamEventsSimple(team, apiKey string, opts *RequestOptions) ([]responses.EventSimple, int, error) {
	// Generate the request
	req := newRequest("/team/"+team+"/events/simple", apiKey, opts)

	// Send the request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, 0, err
	}

	// Check the response type
	if resp.StatusCode == http.StatusUnauthorized {
		return nil, resp.StatusCode, ErrInvalidToken
	} else if resp.StatusCode == http.StatusNotModified {
		return nil, resp.StatusCode, nil
	} else if resp.StatusCode != 200 {
		return nil, resp.StatusCode, err
	}

	// Decode body
	var events []responses.EventSimple
	if err := json.NewDecoder(resp.Body).Decode(&events); err != nil {
		return nil, resp.StatusCode, err
	}

	return events, resp.StatusCode, nil
}

// Get a list of keys for events a team is in
func TeamEventsKey(team, apiKey string, opts *RequestOptions) ([]string, int, error) {
	// Generate the request
	req := newRequest("/team/"+team+"/events/keys", apiKey, opts)

	// Send the request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, 0, err
	}

	// Check the response type
	if resp.StatusCode == http.StatusUnauthorized {
		return nil, resp.StatusCode, ErrInvalidToken
	} else if resp.StatusCode == http.StatusNotModified {
		return nil, resp.StatusCode, nil
	} else if resp.StatusCode != 200 {
		return nil, resp.StatusCode, err
	}

	// Decode body
	var events []string
	if err := json.NewDecoder(resp.Body).Decode(&events); err != nil {
		return nil, resp.StatusCode, err
	}

	return events, resp.StatusCode, nil
}

// Get a list of events a team is in for a given year
func TeamEventsByYear(team string, year int64, apiKey string, opts *RequestOptions) ([]responses.Event, int, error) {
	// Generate the request
	req := newRequest("/team/"+team+"/events/"+strconv.FormatInt(year, 10), apiKey, opts)

	// Send the request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, 0, err
	}

	// Check the response type
	if resp.StatusCode == http.StatusUnauthorized {
		return nil, resp.StatusCode, ErrInvalidToken
	} else if resp.StatusCode == http.StatusNotModified {
		return nil, resp.StatusCode, nil
	} else if resp.StatusCode != 200 {
		return nil, resp.StatusCode, err
	}

	// Decode body
	var events []responses.Event
	if err := json.NewDecoder(resp.Body).Decode(&events); err != nil {
		return nil, resp.StatusCode, err
	}

	return events, resp.StatusCode, nil
}

// Get a simplified list of events a team is in for a given year
func TeamEventsByYearSimple(team string, year int64, apiKey string, opts *RequestOptions) ([]responses.EventSimple, int, error) {
	// Generate the request
	req := newRequest("/team/"+team+"/events/"+strconv.FormatInt(year, 10)+"/simple", apiKey, opts)

	// Send the request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, 0, err
	}

	// Check the response type
	if resp.StatusCode == http.StatusUnauthorized {
		return nil, resp.StatusCode, ErrInvalidToken
	} else if resp.StatusCode == http.StatusNotModified {
		return nil, resp.StatusCode, nil
	} else if resp.StatusCode != 200 {
		return nil, resp.StatusCode, err
	}

	// Decode body
	var events []responses.EventSimple
	if err := json.NewDecoder(resp.Body).Decode(&events); err != nil {
		return nil, resp.StatusCode, err
	}

	return events, resp.StatusCode, nil
}

// Get a list of keys for events a team is in for a given year
func TeamEventsByYearKey(team string, year int64, apiKey string, opts *RequestOptions) ([]string, int, error) {
	// Generate the request
	req := newRequest("/team/"+team+"/events/"+strconv.FormatInt(year, 10)+"/keys", apiKey, opts)

	// Send the request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, 0, err
	}

	// Check the response type
	if resp.StatusCode == http.StatusUnauthorized {
		return nil, resp.StatusCode, ErrInvalidToken
	} else if resp.StatusCode == http.StatusNotModified {
		return nil, resp.StatusCode, nil
	} else if resp.StatusCode != 200 {
		return nil, resp.StatusCode, err
	}

	// Decode body
	var events []string
	if err := json.NewDecoder(resp.Body).Decode(&events); err != nil {
		return nil, resp.StatusCode, err
	}

	return events, resp.StatusCode, nil
}

// Get a map of a team's statuses for events in a given year
func TeamEventStatusesByYear(team string, year int64, apiKey string, opts *RequestOptions) (map[string]responses.TeamEventStatus, int, error) {
	// Generate the request
	req := newRequest("/team/"+team+"/events/"+strconv.FormatInt(year, 10)+"/statuses", apiKey, opts)

	// Send the request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, 0, err
	}

	// Check the response type
	if resp.StatusCode == http.StatusUnauthorized {
		return nil, resp.StatusCode, ErrInvalidToken
	} else if resp.StatusCode == http.StatusNotModified {
		return nil, resp.StatusCode, nil
	} else if resp.StatusCode != 200 {
		return nil, resp.StatusCode, err
	}

	// Decode body
	var statuses map[string]responses.TeamEventStatus
	if err := json.NewDecoder(resp.Body).Decode(&statuses); err != nil {
		return nil, resp.StatusCode, err
	}

	return statuses, resp.StatusCode, nil
}

// Get a list of matches a team is in for an event
func TeamEventMatches(team, event, apiKey string, opts *RequestOptions) ([]responses.Match, int, error) {
	// Generate the request
	req := newRequest("/team/"+team+"/event/"+event+"/matches", apiKey, opts)

	// Send the request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, 0, err
	}

	// Check the response type
	if resp.StatusCode == http.StatusUnauthorized {
		return nil, resp.StatusCode, ErrInvalidToken
	} else if resp.StatusCode == http.StatusNotModified {
		return nil, resp.StatusCode, nil
	} else if resp.StatusCode != 200 {
		return nil, resp.StatusCode, err
	}

	// Decode body
	var matches []responses.Match
	if err := json.NewDecoder(resp.Body).Decode(&matches); err != nil {
		return nil, resp.StatusCode, err
	}

	return matches, resp.StatusCode, nil
}

// Get a simplified list of matches a team is in for an event
func TeamEventMatchesSimple(team, event, apiKey string, opts *RequestOptions) ([]responses.MatchSimple, int, error) {
	// Generate the request
	req := newRequest("/team/"+team+"/event/"+event+"/matches/simple", apiKey, opts)

	// Send the request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, 0, err
	}

	// Check the response type
	if resp.StatusCode == http.StatusUnauthorized {
		return nil, resp.StatusCode, ErrInvalidToken
	} else if resp.StatusCode == http.StatusNotModified {
		return nil, resp.StatusCode, nil
	} else if resp.StatusCode != 200 {
		return nil, resp.StatusCode, err
	}

	// Decode body
	var matches []responses.MatchSimple
	if err := json.NewDecoder(resp.Body).Decode(&matches); err != nil {
		return nil, resp.StatusCode, err
	}

	return matches, resp.StatusCode, nil
}

// Get a simplified list of matches a team is in for an event
func TeamEventMatchesKey(team, event, apiKey string, opts *RequestOptions) ([]string, int, error) {
	// Generate the request
	req := newRequest("/team/"+team+"/event/"+event+"/matches/keys", apiKey, opts)

	// Send the request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, 0, err
	}

	// Check the response type
	if resp.StatusCode == http.StatusUnauthorized {
		return nil, resp.StatusCode, ErrInvalidToken
	} else if resp.StatusCode == http.StatusNotModified {
		return nil, resp.StatusCode, nil
	} else if resp.StatusCode != 200 {
		return nil, resp.StatusCode, err
	}

	// Decode body
	var matches []string
	if err := json.NewDecoder(resp.Body).Decode(&matches); err != nil {
		return nil, resp.StatusCode, err
	}

	return matches, resp.StatusCode, nil
}

// Get the awards for a team at an event
func TeamEventAwards(team, event, apiKey string, opts *RequestOptions) ([]responses.Award, int, error) {
	// Generate the request
	req := newRequest("/team/"+team+"/event/"+event+"/awards", apiKey, opts)

	// Send the request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, 0, err
	}

	// Check the response type
	if resp.StatusCode == http.StatusUnauthorized {
		return nil, resp.StatusCode, ErrInvalidToken
	} else if resp.StatusCode == http.StatusNotModified {
		return nil, resp.StatusCode, nil
	} else if resp.StatusCode != 200 {
		return nil, resp.StatusCode, err
	}

	// Decode body
	var awards []responses.Award
	if err := json.NewDecoder(resp.Body).Decode(&awards); err != nil {
		return nil, resp.StatusCode, err
	}

	return awards, resp.StatusCode, nil
}

// Get the status of a team at an event
func TeamEventStatus(team, event, apiKey string, opts *RequestOptions) (*responses.TeamEventStatus, int, error) {
	// Generate the request
	req := newRequest("/team/"+team+"/event/"+event+"/status", apiKey, opts)

	// Send the request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, 0, err
	}

	// Check the response type
	if resp.StatusCode == http.StatusUnauthorized {
		return nil, resp.StatusCode, ErrInvalidToken
	} else if resp.StatusCode == http.StatusNotModified {
		return nil, resp.StatusCode, nil
	} else if resp.StatusCode != 200 {
		return nil, resp.StatusCode, err
	}

	// Decode body
	var status responses.TeamEventStatus
	if err := json.NewDecoder(resp.Body).Decode(&status); err != nil {
		return nil, resp.StatusCode, err
	}

	return &status, resp.StatusCode, nil
}

// Get all the awards a team has received
func TeamAwards(team, apiKey string, opts *RequestOptions) ([]responses.Award, int, error) {
	// Generate the request
	req := newRequest("/team/"+team+"/awards", apiKey, opts)

	// Send the request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, 0, err
	}

	// Check the response type
	if resp.StatusCode == http.StatusUnauthorized {
		return nil, resp.StatusCode, ErrInvalidToken
	} else if resp.StatusCode == http.StatusNotModified {
		return nil, resp.StatusCode, nil
	} else if resp.StatusCode != 200 {
		return nil, resp.StatusCode, err
	}

	// Decode body
	var awards []responses.Award
	if err := json.NewDecoder(resp.Body).Decode(&awards); err != nil {
		return nil, resp.StatusCode, err
	}

	return awards, resp.StatusCode, nil
}

// Get a list of awards given to a team in a given year
func TeamAwardsByYear(team string, year int64, apiKey string, opts *RequestOptions) ([]responses.Award, int, error) {
	// Generate the request
	req := newRequest("/team/"+team+"/"+strconv.FormatInt(year, 10)+"/awards", apiKey, opts)

	// Send the request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, 0, err
	}

	// Check the response type
	if resp.StatusCode == http.StatusUnauthorized {
		return nil, resp.StatusCode, ErrInvalidToken
	} else if resp.StatusCode == http.StatusNotModified {
		return nil, resp.StatusCode, nil
	} else if resp.StatusCode != 200 {
		return nil, resp.StatusCode, err
	}

	// Decode body
	var awards []responses.Award
	if err := json.NewDecoder(resp.Body).Decode(&awards); err != nil {
		return nil, resp.StatusCode, err
	}

	return awards, resp.StatusCode, nil
}

// Get the matches a team was in for a given year
func TeamMatchesByYear(team string, year int64, apiKey string, opts *RequestOptions) ([]responses.Match, int, error) {
	// Generate the request
	req := newRequest("/team/"+team+"/matches/"+strconv.FormatInt(year, 10), apiKey, opts)

	// Send the request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, 0, err
	}

	// Check the response type
	if resp.StatusCode == http.StatusUnauthorized {
		return nil, resp.StatusCode, ErrInvalidToken
	} else if resp.StatusCode == http.StatusNotModified {
		return nil, resp.StatusCode, nil
	} else if resp.StatusCode != 200 {
		return nil, resp.StatusCode, err
	}

	// Decode body
	var matches []responses.Match
	if err := json.NewDecoder(resp.Body).Decode(&matches); err != nil {
		return nil, resp.StatusCode, err
	}

	return matches, resp.StatusCode, nil
}

// Get a simplified list of matches a team was in for a given year
func TeamMatchesByYearSimple(team string, year int64, apiKey string, opts *RequestOptions) ([]responses.MatchSimple, int, error) {
	// Generate the request
	req := newRequest("/team/"+team+"/matches/"+strconv.FormatInt(year, 10), apiKey, opts)

	// Send the request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, 0, err
	}

	// Check the response type
	if resp.StatusCode == http.StatusUnauthorized {
		return nil, resp.StatusCode, ErrInvalidToken
	} else if resp.StatusCode == http.StatusNotModified {
		return nil, resp.StatusCode, nil
	} else if resp.StatusCode != 200 {
		return nil, resp.StatusCode, err
	}

	// Decode body
	var matches []responses.MatchSimple
	if err := json.NewDecoder(resp.Body).Decode(&matches); err != nil {
		return nil, resp.StatusCode, err
	}

	return matches, resp.StatusCode, nil
}

// Get the keys of the matches a team was in for a given year
func TeamMatchesByYearKey(team string, year int64, apiKey string, opts *RequestOptions) ([]string, int, error) {
	// Generate the request
	req := newRequest("/team/"+team+"/matches/"+strconv.FormatInt(year, 10), apiKey, opts)

	// Send the request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, 0, err
	}

	// Check the response type
	if resp.StatusCode == http.StatusUnauthorized {
		return nil, resp.StatusCode, ErrInvalidToken
	} else if resp.StatusCode == http.StatusNotModified {
		return nil, resp.StatusCode, nil
	} else if resp.StatusCode != 200 {
		return nil, resp.StatusCode, err
	}

	// Decode body
	var matches []string
	if err := json.NewDecoder(resp.Body).Decode(&matches); err != nil {
		return nil, resp.StatusCode, err
	}

	return matches, resp.StatusCode, nil
}

// Get the media of a team for a given year
func TeamMediaByYear(team string, year int64, apiKey string, opts *RequestOptions) ([]responses.Media, int, error) {
	// Generate the request
	req := newRequest("/team/"+team+"/media/"+strconv.FormatInt(year, 10), apiKey, opts)

	// Send the request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, 0, err
	}

	// Check the response type
	if resp.StatusCode == http.StatusUnauthorized {
		return nil, resp.StatusCode, ErrInvalidToken
	} else if resp.StatusCode == http.StatusNotModified {
		return nil, resp.StatusCode, nil
	} else if resp.StatusCode != 200 {
		return nil, resp.StatusCode, err
	}

	// Decode body
	var media []responses.Media
	if err := json.NewDecoder(resp.Body).Decode(&media); err != nil {
		return nil, resp.StatusCode, err
	}

	return media, resp.StatusCode, nil
}

// Get a list of media for a given team with a given tag
func TeamMediaByTag(team, tag, apiKey string, opts *RequestOptions) ([]responses.Media, int, error) {
	// Generate the request
	req := newRequest("/team/"+team+"/media/tag/"+tag, apiKey, opts)

	// Send the request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, 0, err
	}

	// Check the response type
	if resp.StatusCode == http.StatusUnauthorized {
		return nil, resp.StatusCode, ErrInvalidToken
	} else if resp.StatusCode == http.StatusNotModified {
		return nil, resp.StatusCode, nil
	} else if resp.StatusCode != 200 {
		return nil, resp.StatusCode, err
	}

	// Decode body
	var media []responses.Media
	if err := json.NewDecoder(resp.Body).Decode(&media); err != nil {
		return nil, resp.StatusCode, err
	}

	return media, resp.StatusCode, nil
}

// Get a list of media for a given team with a given tag and year
func TeamMediaByTagAndYear(team, tag string, year int64, apiKey string, opts *RequestOptions) ([]responses.Media, int, error) {
	// Generate the request
	req := newRequest("/team/"+team+"/media/tag/"+tag+"/"+strconv.FormatInt(year, 10), apiKey, opts)

	// Send the request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, 0, err
	}

	// Check the response type
	if resp.StatusCode == http.StatusUnauthorized {
		return nil, resp.StatusCode, ErrInvalidToken
	} else if resp.StatusCode == http.StatusNotModified {
		return nil, resp.StatusCode, nil
	} else if resp.StatusCode != 200 {
		return nil, resp.StatusCode, err
	}

	// Decode body
	var media []responses.Media
	if err := json.NewDecoder(resp.Body).Decode(&media); err != nil {
		return nil, resp.StatusCode, err
	}

	return media, resp.StatusCode, nil
}

// Get a list of social media for a team
func TeamSocialMedia(team, apiKey string, opts *RequestOptions) ([]responses.Media, int, error) {
	// Generate the request
	req := newRequest("/team/"+team+"/social_media", apiKey, opts)

	// Send the request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, 0, err
	}

	// Check the response type
	if resp.StatusCode == http.StatusUnauthorized {
		return nil, resp.StatusCode, ErrInvalidToken
	} else if resp.StatusCode == http.StatusNotModified {
		return nil, resp.StatusCode, nil
	} else if resp.StatusCode != 200 {
		return nil, resp.StatusCode, err
	}

	// Decode body
	var socialMedia []responses.Media
	if err := json.NewDecoder(resp.Body).Decode(&socialMedia); err != nil {
		return nil, resp.StatusCode, err
	}

	return socialMedia, resp.StatusCode, nil
}
