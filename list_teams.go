package tba

import (
	"encoding/json"
	"github.com/akrantz01/go-tba/responses"
	"net/http"
	"strconv"
)

// Get a list of teams in the TBA database.
// The response is paginated by team numbers in sets of 500.
// The set of 500 is not number of teams, but the range of team numbers.
func ListTeams(page int64, apiKey string, opts *RequestOptions) ([]responses.Team, int, error) {
	// Generate the request
	req := newRequest("/teams/"+strconv.FormatInt(page, 10), apiKey, opts)

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

// Get a list of teams with simplified data in the TBA database.
// The response is paginated by team numbers in sets of 500.
// The set of 500 is not number of teams, but the range of team numbers.
func ListTeamsSimple(page int64, apiKey string, opts *RequestOptions) ([]responses.TeamSimple, int, error) {
	// Generate the request
	req := newRequest("/teams/"+strconv.FormatInt(page, 10)+"/simple", apiKey, opts)

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

// Get a list of team keys in the TBA database.
// The response is paginated by team numbers in sets of 500.
// The set of 500 is not number of teams, but the range of team numbers.
func ListTeamKeys(page int64, apiKey string, opts *RequestOptions) ([]string, int, error) {
	// Generate the request
	req := newRequest("/teams/"+strconv.FormatInt(page, 10)+"/keys", apiKey, opts)

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

// Get a list of teams in the TBA database for a specific year.
// The response is paginated by team numbers in sets of 500.
// The set of 500 is not number of teams, but the range of team numbers.
func ListTeamsByYear(year, page int64, apiKey string, opts *RequestOptions) ([]responses.Team, int, error) {
	// Generate the request
	req := newRequest("/teams/"+strconv.FormatInt(year, 10)+"/"+strconv.FormatInt(page, 10), apiKey, opts)

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

// Get a list of teams with simplified data in the TBA database for a specific year.
// The response is paginated by team numbers in sets of 500.
// The set of 500 is not number of teams, but the range of team numbers.
func ListTeamsByYearSimple(year, page int64, apiKey string, opts *RequestOptions) ([]responses.TeamSimple, int, error) {
	// Generate the request
	req := newRequest("/teams/"+strconv.FormatInt(year, 10)+"/"+strconv.FormatInt(page, 10)+"/simple", apiKey, opts)

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

// Get a list of team keys in the TBA database for a specific year.
// The response is paginated by team numbers in sets of 500.
// The set of 500 is not number of teams, but the range of team numbers.
func ListTeamsByYearKey(year, page int64, apiKey string, opts *RequestOptions) ([]string, int, error) {
	// Generate the request
	req := newRequest("/teams/"+strconv.FormatInt(year, 10)+"/"+strconv.FormatInt(page, 10)+"/keys", apiKey, opts)

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
