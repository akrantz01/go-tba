package tba

import (
	"encoding/json"
	"github.com/akrantz01/go-tba/responses"
	"net/http"
)

// Get information about a match
func Match(key, apiKey string, opts *RequestOptions) (*responses.Match, int, error) {
	// Generate the request
	req := newRequest("/matches/"+key, apiKey, opts)

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

	return &match, resp.StatusCode, nil
}

// Get a simplified information about a match
func MatchSimple(key, apiKey string, opts *RequestOptions) (*responses.MatchSimple, int, error) {
	// Generate the request
	req := newRequest("/matches/"+key+"/simple", apiKey, opts)

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
	req := newRequest("/matches/"+key+"/timeseries", apiKey, opts)

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
