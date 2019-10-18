package tba

import (
	"encoding/json"
	"github.com/akrantz01/go-tba/responses"
	"net/http"
)

// Get a list of team rankings in a district
func ListDistrictRankings(district, apiKey string, opts *RequestOptions) ([]responses.DistrictRanking, int, error) {
	// Generate the request
	req := newRequest("/district/"+district+"/rankings", apiKey, opts)

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
	var districtRankings []responses.DistrictRanking
	if err := json.NewDecoder(resp.Body).Decode(&districtRankings); err != nil {
		return nil, resp.StatusCode, err
	}

	return districtRankings, resp.StatusCode, nil
}

// Get a list of team rankings for a given event
func ListEventRankings(event, apiKey string, opts *RequestOptions) ([]responses.EventRanking, int, error) {
	// Generate the request
	req := newRequest("/event/"+event+"/rankings", apiKey, opts)

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
	var eventRankings []responses.EventRanking
	if err := json.NewDecoder(resp.Body).Decode(&eventRankings); err != nil {
		return nil, resp.StatusCode, err
	}

	return eventRankings, resp.StatusCode, nil
}
