package tba

import (
	"encoding/json"
	"github.com/akrantz01/go-tba/responses"
	"net/http"
)

// Get the statuses of the teams at a given event
func ListTeamStatusesAtEvent(event, apiKey string, opts *RequestOptions) ([]responses.TeamEventStatus, int, error) {
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

	return teamEventStatuses, resp.StatusCode, nil
}
