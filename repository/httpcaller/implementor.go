package httpcaller

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/LukmanulHakim18/time2go/model"
)

type HttpCallerClient struct{}

func (c *HttpCallerClient) HealthCheck(ctx context.Context) error {
	return nil
}

func (h *HttpCallerClient) ExecuteEvent(ctx context.Context, e model.HTTPRequestConfig) (*http.Response, error) {
	req, err := h.buildRequest(e)
	if err != nil {
		return nil, err
	}
	c := &http.Client{
		Timeout: e.Timeout,
	}
	return c.Do(req)
}

func (*HttpCallerClient) buildRequest(e model.HTTPRequestConfig) (*http.Request, error) {
	u, err := url.Parse(e.URL)
	if err != nil {
		return nil, fmt.Errorf("invalid URL: %w", err)
	}

	// Attach query parameters
	q := u.Query()
	for key, val := range e.QueryParams {
		q.Set(key, val)
	}
	u.RawQuery = q.Encode()

	var bodyBytes []byte
	if e.Body != nil {
		bodyBytes, err = json.Marshal(e.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal body: %w", err)
		}
	}

	req, err := http.NewRequest(e.Method, u.String(), bytes.NewBuffer(bodyBytes))
	if err != nil {
		return nil, fmt.Errorf("failed to build request: %w", err)
	}

	// Set headers
	for k, v := range e.Headers {
		req.Header.Set(k, v)
	}
	if e.Body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	return req, nil
}
