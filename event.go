package tba

import (
	"encoding/json"
	"github.com/akrantz01/go-tba/responses"
	"github.com/mitchellh/mapstructure"
	"net/http"
	"strconv"
)

// Get a single event by its key
func Event(key, apiKey string, opts *RequestOptions) (*responses.Event, int, error) {
	// Generate the request
	req := newRequest("/event/"+key, apiKey, opts)

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
	var event responses.Event
	if err := json.NewDecoder(resp.Body).Decode(&event); err != nil {
		return nil, resp.StatusCode, err
	}

	// Close the body
	if err := resp.Body.Close(); err != nil {
		return nil, resp.StatusCode, err
	}

	return &event, resp.StatusCode, nil
}

// Get the simple representation of a single event by its key
func EventSimple(key, apiKey string, opts *RequestOptions) (*responses.EventSimple, int, error) {
	// Generate the request
	req := newRequest("/event/"+key+"/simple", apiKey, opts)

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
	var event responses.EventSimple
	if err := json.NewDecoder(resp.Body).Decode(&event); err != nil {
		return nil, resp.StatusCode, err
	}

	// Close the body
	if err := resp.Body.Close(); err != nil {
		return nil, resp.StatusCode, err
	}

	return &event, resp.StatusCode, nil
}

// Get the alliances in an event
func EventAlliances(key, apiKey string, opts *RequestOptions) ([]responses.EventAlliance, int, error) {
	// Generate the request
	req := newRequest("/event/"+key+"/alliances", apiKey, opts)

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
	var alliances []responses.EventAlliance
	if err := json.NewDecoder(resp.Body).Decode(&alliances); err != nil {
		return nil, resp.StatusCode, err
	}

	// Close the body
	if err := resp.Body.Close(); err != nil {
		return nil, resp.StatusCode, err
	}

	return alliances, resp.StatusCode, nil
}

// Get the insights for an event.
// The year must be between 2016 and 2019, inclusive.
func EventInsights(key string, year int64, apiKey string, opts *RequestOptions) (interface{}, int, error) {
	// Ensure year in valid range for insights
	if year < 2016 || year > 2019 {
		return nil, 0, ErrYearOutOfBounds
	}

	// Generate the request
	req := newRequest("/event/"+key+"/insights", apiKey, opts)

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

	// Conditionally decode body based on year
	switch year {
	case 2016:
		var insights responses.EventInsights2016
		if err := json.NewDecoder(resp.Body).Decode(&insights); err != nil {
			return nil, resp.StatusCode, err
		}
		// Close the body
		if err := resp.Body.Close(); err != nil {
			return nil, resp.StatusCode, err
		}
		return insights, resp.StatusCode, nil

	case 2017:
		var insights responses.EventInsights2016
		if err := json.NewDecoder(resp.Body).Decode(&insights); err != nil {
			return nil, resp.StatusCode, err
		}
		// Close the body
		if err := resp.Body.Close(); err != nil {
			return nil, resp.StatusCode, err
		}
		return insights, resp.StatusCode, nil

	case 2018:
		var insights responses.EventInsights2016
		if err := json.NewDecoder(resp.Body).Decode(&insights); err != nil {
			return nil, resp.StatusCode, err
		}
		// Close the body
		if err := resp.Body.Close(); err != nil {
			return nil, resp.StatusCode, err
		}
		return insights, resp.StatusCode, nil

	case 2019:
		var insights responses.EventInsights2016
		if err := json.NewDecoder(resp.Body).Decode(&insights); err != nil {
			return nil, resp.StatusCode, err
		}
		// Close the body
		if err := resp.Body.Close(); err != nil {
			return nil, resp.StatusCode, err
		}
		return insights, resp.StatusCode, nil

	default:
		// This shouldn't be reachable
		return nil, 0, ErrYearOutOfBounds
	}
}

// Get OPRs, DPRs, and CCWMs for a given event
func EventOPRs(key, apiKey string, opts *RequestOptions) (*responses.EventOPRs, int, error) {
	// Generate the request
	req := newRequest("/event/"+key+"/oprs", apiKey, opts)

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
	var oprs responses.EventOPRs
	if err := json.NewDecoder(resp.Body).Decode(&oprs); err != nil {
		return nil, resp.StatusCode, err
	}

	// Close the body
	if err := resp.Body.Close(); err != nil {
		return nil, resp.StatusCode, err
	}

	return &oprs, resp.StatusCode, nil
}

// Get predictions for a given event
func EventPredictions(key, apiKey string, opts *RequestOptions) (*responses.EventPredictions, int, error) {
	// Generate the request
	req := newRequest("/event/"+key+"/predictions", apiKey, opts)

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
	var predictions responses.EventPredictions
	if err := json.NewDecoder(resp.Body).Decode(&predictions); err != nil {
		return nil, resp.StatusCode, err
	}

	// Close the body
	if err := resp.Body.Close(); err != nil {
		return nil, resp.StatusCode, err
	}

	return &predictions, resp.StatusCode, nil
}

// Get rankings for a given event
func EventRankings(key, apiKey string, opts *RequestOptions) ([]responses.EventRanking, int, error) {
	// Generate the request
	req := newRequest("/event/"+key+"/rankings", apiKey, opts)

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
	var rankings []responses.EventRanking
	if err := json.NewDecoder(resp.Body).Decode(&rankings); err != nil {
		return nil, resp.StatusCode, err
	}

	// Close the body
	if err := resp.Body.Close(); err != nil {
		return nil, resp.StatusCode, err
	}

	return rankings, resp.StatusCode, nil
}

// Get district points for a given event
func EventDistrictPoints(key, apiKey string, opts *RequestOptions) (*responses.EventDistrictPoints, int, error) {
	// Generate the request
	req := newRequest("/event/"+key+"/district_points", apiKey, opts)

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
	var districtPoints responses.EventDistrictPoints
	if err := json.NewDecoder(resp.Body).Decode(&districtPoints); err != nil {
		return nil, resp.StatusCode, err
	}

	// Close the body
	if err := resp.Body.Close(); err != nil {
		return nil, resp.StatusCode, err
	}

	return &districtPoints, resp.StatusCode, nil
}

// Get teams in an event
func EventTeams(key, apiKey string, opts *RequestOptions) ([]responses.Team, int, error) {
	// Generate the request
	req := newRequest("/event/"+key+"/teams", apiKey, opts)

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
	var teams []responses.Team
	if err := json.NewDecoder(resp.Body).Decode(&teams); err != nil {
		return nil, resp.StatusCode, err
	}

	// Close the body
	if err := resp.Body.Close(); err != nil {
		return nil, resp.StatusCode, err
	}

	return teams, resp.StatusCode, nil
}

// Get teams in an event
func EventTeamsSimple(key, apiKey string, opts *RequestOptions) ([]responses.TeamSimple, int, error) {
	// Generate the request
	req := newRequest("/event/"+key+"/teams/simple", apiKey, opts)

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
	var teams []responses.TeamSimple
	if err := json.NewDecoder(resp.Body).Decode(&teams); err != nil {
		return nil, resp.StatusCode, err
	}

	// Close the body
	if err := resp.Body.Close(); err != nil {
		return nil, resp.StatusCode, err
	}

	return teams, resp.StatusCode, nil
}

// Get teams in an event
func EventTeamKeys(key, apiKey string, opts *RequestOptions) ([]string, int, error) {
	// Generate the request
	req := newRequest("/event/"+key+"/teams/key", apiKey, opts)

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
	var teams []string
	if err := json.NewDecoder(resp.Body).Decode(&teams); err != nil {
		return nil, resp.StatusCode, err
	}

	// Close the body
	if err := resp.Body.Close(); err != nil {
		return nil, resp.StatusCode, err
	}

	return teams, resp.StatusCode, nil
}

// Get the statuses of the teams at a given event
func EventTeamStatuses(event, apiKey string, opts *RequestOptions) ([]responses.TeamEventStatus, int, error) {
	// Generate the request
	req := newRequest("/event/"+event+"/teams/statuses", apiKey, opts)

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
	var teamEventStatuses []responses.TeamEventStatus
	if err := json.NewDecoder(resp.Body).Decode(&teamEventStatuses); err != nil {
		return nil, resp.StatusCode, err
	}

	// Close the body
	if err := resp.Body.Close(); err != nil {
		return nil, resp.StatusCode, err
	}

	return teamEventStatuses, resp.StatusCode, nil
}

// Get the matches of an event
func EventMatches(event, apiKey string, opts *RequestOptions) ([]responses.Match, int, error) {
	// Generate the request
	req := newRequest("/event/"+event+"/matches", apiKey, opts)

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

	// Close the body
	if err := resp.Body.Close(); err != nil {
		return nil, resp.StatusCode, err
	}

	for _, match := range matches {
		// Coerce score breakdown type
		year, _ := strconv.ParseInt(match.Key[:4], 10, 64)
		switch year {
		case 2019:
			var score responses.ScoringBreakdown2019
			if err := mapstructure.Decode(match.ScoreBreakdown, &score); err != nil {
				return nil, resp.StatusCode, err
			}
			match.ScoreBreakdown = score

		case 2018:
			var score responses.ScoringBreakdown2018
			if err := mapstructure.Decode(match.ScoreBreakdown, &score); err != nil {
				return nil, resp.StatusCode, err
			}
			match.ScoreBreakdown = score

		case 2017:
			var score responses.ScoringBreakdown2017
			if err := mapstructure.Decode(match.ScoreBreakdown, &score); err != nil {
				return nil, resp.StatusCode, err
			}
			match.ScoreBreakdown = score

		case 2016:
			var score responses.ScoringBreakdown2016
			if err := mapstructure.Decode(match.ScoreBreakdown, &score); err != nil {
				return nil, resp.StatusCode, err
			}
			match.ScoreBreakdown = score
		}
	}

	return matches, resp.StatusCode, nil
}

// Get a simple representation of the matches of an event
func EventMatchesSimple(event, apiKey string, opts *RequestOptions) ([]responses.MatchSimple, int, error) {
	// Generate the request
	req := newRequest("/event/"+event+"/matches/simple", apiKey, opts)

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

	// Close the body
	if err := resp.Body.Close(); err != nil {
		return nil, resp.StatusCode, err
	}

	return matches, resp.StatusCode, nil
}

// Get the keys of the matches of an event
func EventMatchKeys(event, apiKey string, opts *RequestOptions) ([]string, int, error) {
	// Generate the request
	req := newRequest("/event/"+event+"/matches/key", apiKey, opts)

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

	// Close the body
	if err := resp.Body.Close(); err != nil {
		return nil, resp.StatusCode, err
	}

	return matches, resp.StatusCode, nil
}

// Get matches with timeseries data for an event
func EventTimeseries(event, apiKey string, opts *RequestOptions) ([]string, int, error) {
	// Generate the request
	req := newRequest("/event/"+event+"/matches/timeseries", apiKey, opts)

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
	var timeseries []string
	if err := json.NewDecoder(resp.Body).Decode(&timeseries); err != nil {
		return nil, resp.StatusCode, err
	}

	// Close the body
	if err := resp.Body.Close(); err != nil {
		return nil, resp.StatusCode, err
	}

	return timeseries, resp.StatusCode, nil
}

// Get awards given at an event
func EventAwards(event, apiKey string, opts *RequestOptions) ([]responses.Award, int, error) {
	// Generate the request
	req := newRequest("/event/"+event+"/awards", apiKey, opts)

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

	// Close the body
	if err := resp.Body.Close(); err != nil {
		return nil, resp.StatusCode, err
	}

	return awards, resp.StatusCode, nil
}
