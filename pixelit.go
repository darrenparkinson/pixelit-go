// Copyright 2023 Darren Parkinson. All rights reserved.
//
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

//go:generate go run gen-accessors.go

package pixelit

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

// Client is the main pixelit client for interacting with the library.  It can be created using NewClient
type Client struct {
	// BaseURL for PixelIt Display.  Set using `pixelit.NewClient()` automatically via the host parameter, or set directly.
	BaseURL string

	// Host IP for PixelIt Display.  Set using `pixelit.NewClient()`, or set directly.
	Host string

	//HTTP Client to use for making requests, allowing the user to supply their own if required.
	HTTPClient *http.Client
}

// NewClient is a helper function that returns a new pixelit client given a host name or IP.
// Optionally you can provide your own http client or use nil to use the default.  This is done to
// ensure you're aware of the decision you're making to not provide your own http client.
func NewClient(host string, client *http.Client) (*Client, error) {
	if host == "" {
		return nil, errors.New("host required")
	}
	if client == nil {
		client = &http.Client{
			Timeout: 10 * time.Second,
		}
	}
	c := &Client{
		BaseURL:    fmt.Sprintf("http://%s/api", host),
		Host:       host,
		HTTPClient: client,
	}
	return c, nil
}

// Bool is a helper routine that allocates a new bool value
// to store v and returns a pointer to it.
func Bool(v bool) *bool { return &v }

// Int is a helper routine that allocates a new int value
// to store v and returns a pointer to it.
func Int(v int) *int { return &v }

// Int64 is a helper routine that allocates a new int64 value
// to store v and returns a pointer to it.
func Int64(v int64) *int64 { return &v }

// Float64 is a helper routine that allocates a new Float64 value
// to store v and returns a pointer to it.
func Float64(v float64) *float64 { return &v }

// String is a helper routine that allocates a new string value
// to store v and returns a pointer to it.
func String(v string) *string { return &v }

// makeRequest provides a single function to add common items to the request.
func (c *Client) makeRequest(ctx context.Context, req *http.Request, v interface{}) error {
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	rc := req.WithContext(ctx)
	res, err := c.HTTPClient.Do(rc)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {

		var pixelitErr error

		switch res.StatusCode {
		case 400:
			pixelitErr = ErrBadRequest
		case 401:
			pixelitErr = ErrUnauthorized
		case 403:
			pixelitErr = ErrForbidden
		case 404:
			pixelitErr = ErrNotFound
		case 500:
			pixelitErr = ErrInternalError
		default:
			pixelitErr = ErrUnknown
		}

		return pixelitErr

	}

	if res.StatusCode == http.StatusCreated {
		return nil
	}

	if err = json.NewDecoder(res.Body).Decode(&v); err != nil {
		return err
	}

	return nil
}
