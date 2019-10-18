package tba

import (
	"encoding/json"
	"github.com/akrantz01/go-tba/responses"
	"net/http"
)

// Get the status of the TBA API and internal status information
func Status(apiKey string, opts *RequestOptions) (*responses.ApiStatus, int, error) {
	// Generate the request
	req := newRequest("/status", apiKey, opts)

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
	var status responses.ApiStatus
	if err := json.NewDecoder(resp.Body).Decode(&status); err != nil {
		return nil, resp.StatusCode, err
	}

	return &status, resp.StatusCode, nil
}
