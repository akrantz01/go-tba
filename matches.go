package tba

import (
	"encoding/json"
	"github.com/akrantz01/go-tba/responses"
	"github.com/mitchellh/mapstructure"
	"net/http"
	"strconv"
)

// Get information about a match
func Match(key, apiKey string, opts *RequestOptions) (*responses.Match, int, error) {
	// Generate the request
	req := newRequest("/match/"+key, apiKey, opts)

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
	var match responses.Match
	if err := json.NewDecoder(resp.Body).Decode(&match); err != nil {
		return nil, resp.StatusCode, err
	}

	// Close the body
	if err := resp.Body.Close(); err != nil {
		return nil, resp.StatusCode, err
	}

	// Coerce score breakdown type
	year, _ := strconv.ParseInt(match.Key[:4], 10, 64)
	switch year {
	case 2020:
		var score responses.ScoringBreakdown2020
		if err := mapstructure.Decode(match.ScoreBreakdown, &score); err != nil {
			return nil, resp.StatusCode, err
		}
		match.ScoreBreakdown = score
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

	return &match, resp.StatusCode, nil
}

// Get a simplified information about a match
func MatchSimple(key, apiKey string, opts *RequestOptions) (*responses.MatchSimple, int, error) {
	// Generate the request
	req := newRequest("/match/"+key+"/simple", apiKey, opts)

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
	var match responses.MatchSimple
	if err := json.NewDecoder(resp.Body).Decode(&match); err != nil {
		return nil, resp.StatusCode, err
	}

	// Close the body
	if err := resp.Body.Close(); err != nil {
		return nil, resp.StatusCode, err
	}

	return &match, resp.StatusCode, nil
}

// Get timeseries information about a match
func MatchTimeseries(key, apiKey string, opts *RequestOptions) (*responses.Timeseries2018, int, error) {
	// Generate the request
	req := newRequest("/match/"+key+"/timeseries", apiKey, opts)

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
	var timeseries responses.Timeseries2018
	if err := json.NewDecoder(resp.Body).Decode(&timeseries); err != nil {
		return nil, resp.StatusCode, err
	}

	// Close the body
	if err := resp.Body.Close(); err != nil {
		return nil, resp.StatusCode, err
	}

	return &timeseries, resp.StatusCode, nil
}
