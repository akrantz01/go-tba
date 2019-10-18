package tba

import (
	"encoding/json"
	"github.com/akrantz01/go-tba/responses"
	"net/http"
	"strconv"
)

// Get a list of events for a given year.
func ListEvents(year int64, apiKey string, opts *RequestOptions) ([]responses.Event, int, error) {
	// Generate the request
	req := newRequest("/events/"+strconv.FormatInt(year, 10), apiKey, opts)

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

// Get a simplified list of events for a given year.
func ListEventsSimple(year int64, apiKey string, opts *RequestOptions) ([]responses.EventSimple, int, error) {
	// Generate the request
	req := newRequest("/events/"+strconv.FormatInt(year, 10)+"/simple", apiKey, opts)

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

// Get the keys of events for a given year.
func ListEventsKey(year int64, apiKey string, opts *RequestOptions) ([]string, int, error) {
	// Generate the request
	req := newRequest("/events/"+strconv.FormatInt(year, 10)+"/keys", apiKey, opts)

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

// Get a list of events in a district.
func ListEventsInDistrict(district, apiKey string, opts *RequestOptions) ([]responses.Event, int, error) {
	// Generate the request
	req := newRequest("/district/"+district+"/events", apiKey, opts)

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

// Get a simplified list of events in district
func ListEventsInDistrictSimple(district, apiKey string, opts *RequestOptions) ([]responses.EventSimple, int, error) {
	// Generate the request
	req := newRequest("/district/"+district+"/events/simple", apiKey, opts)

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

// Get a list of the keys for events in a district.
func ListEventsInDistrictKey(district, apiKey string, opts *RequestOptions) ([]string, int, error) {
	// Generate the request
	req := newRequest("/district/"+district+"/events/keys", apiKey, opts)

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
