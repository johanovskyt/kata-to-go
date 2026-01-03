package codewars

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const baseURL = "https://www.codewars.com/api/v1"

type Client struct {
	httpClient *http.Client
}

type Kata struct {
	Name        string `json:"name"`
	URL         string `json:"url"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
}

func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c *Client) GetKata(ctx context.Context, id string) (*Kata, error) {
	reqURL := fmt.Sprintf("%s/code-challenges/%s", baseURL, id)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Accept", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to get response: %w", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var kata Kata
	if err := json.Unmarshal(body, &kata); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &kata, nil
}
