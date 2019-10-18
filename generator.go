package tba

import (
	"errors"
	"net/http"
	"time"
)

// The URL of the TBA v3 read-only API
const tbaApi = "https://www.thebluealliance.com/api/v3"

// Error constants
var ErrInvalidToken = errors.New("invalid api token")
var ErrYearOutOfBounds = errors.New("year out of bounds")

// Additional configuration options for the request being sent
type RequestOptions struct {
	// The value for the Last-Modified header.
	LastModified *time.Time
	// Set a custom user agent header instead of the default Golang user agent.
	UserAgent *string
}

// generateRequest creates a request to retrieve data from the blue alliance.
// This will always be a GET request as we are not pushing new data to The Blue Alliance.
// The authentication header is automatically set.
// Optionally, the Last-Modified and User-Agent headers can be set with the request options.
func newRequest(path, apiKey string, opts *RequestOptions) *http.Request {
	// Build base request
	req, _ := http.NewRequest(http.MethodGet, tbaApi+path, nil)

	// Set authentication header
	req.Header.Set("X-TBA-Auth-Key", apiKey)

	if opts != nil {
		// Set cache control header
		if opts.LastModified != nil {
			req.Header.Set("If-Modified-Since", opts.LastModified.Format("Mon, 02 Jan 2006 15:04:05 GMT"))
		}

		// Set user agent header
		if opts.UserAgent != nil {
			req.Header.Set("User-Agent", *opts.UserAgent)
		}
	}

	return req
}
