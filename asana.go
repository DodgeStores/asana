package asana

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

type Client struct {
	client      *http.Client
	BaseURL     string
	AccessToken string
}

func NewClient(client *http.Client) *Client {
	if client == nil {
		client = &http.Client{}
	}

	return &Client{
		client:      client,
		AccessToken: os.Getenv("ASANA_ACCESS_TOKEN"),
		BaseURL:     "https://app.asana.com/api/1.0",
	}
}

//Request thinly wraps http.Client's NewRequest and Do methods
// while handling the asana access token
func (c *Client) Request(method, uri string, body io.Reader) (*http.Response, error) {

	//Make Sure Base URL is actual url
	_, err := url.Parse(c.BaseURL)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(method, fmt.Sprintf("%s/%s", c.BaseURL, uri), body)

	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.AccessToken))

	return c.client.Do(request)
}
