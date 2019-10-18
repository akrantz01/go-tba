package tba

import (
	"encoding/json"
	"github.com/akrantz01/go-tba/responses"
	"net/http"
	"strconv"
)

// Get the districts for a certain year
func DistrictsByYear(year int64, apiKey string, opts *RequestOptions) ([]responses.DistrictList, int, error) {
	// Generate the request
	req := newRequest("/districts/"+strconv.FormatInt(year, 10), apiKey, opts)

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

// Get events associated with a district
func DistrictEvents(district, apiKey string, opts *RequestOptions) ([]responses.Event, int, error) {
	// Generate the request
	req := newRequest("/districts/"+district+"/events", apiKey, opts)

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

// Get a simplified list of events associated with a district
func DistrictEventsSimple(district, apiKey string, opts *RequestOptions) ([]responses.EventSimple, int, error) {
	// Generate the request
	req := newRequest("/districts/"+district+"/events/simple", apiKey, opts)

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

// Get the keys of the events associated with a district
func DistrictEventsKey(district, apiKey string, opts *RequestOptions) ([]string, int, error) {
	// Generate the request
	req := newRequest("/districts/"+district+"/events/keys", apiKey, opts)

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

// Get the teams associated with a district
func DistrictTeams(district, apiKey string, opts *RequestOptions) ([]responses.Team, int, error) {
	// Generate the request
	req := newRequest("/districts/"+district+"/teams", apiKey, opts)

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

	return teams, resp.StatusCode, nil
}

// Get a simplified list of the teams associated with a district
func DistrictTeamsSimple(district, apiKey string, opts *RequestOptions) ([]responses.TeamSimple, int, error) {
	// Generate the request
	req := newRequest("/districts/"+district+"/teams/simple", apiKey, opts)

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

	return teams, resp.StatusCode, nil
}

// Get the keys of the teams associated with a district
func DistrictTeamsKey(district, apiKey string, opts *RequestOptions) ([]string, int, error) {
	// Generate the request
	req := newRequest("/districts/"+district+"/teams/keys", apiKey, opts)

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

	return teams, resp.StatusCode, nil
}

// Get the rankings of teams in a district
func DistrictRankings(district, apiKey string, opts *RequestOptions) ([]responses.DistrictRanking, int, error) {
	// Generate the request
	req := newRequest("/districts/"+district+"/rankings", apiKey, opts)

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
	var rankings []responses.DistrictRanking
	if err := json.NewDecoder(resp.Body).Decode(&rankings); err != nil {
		return nil, resp.StatusCode, err
	}

	return rankings, resp.StatusCode, nil
}
