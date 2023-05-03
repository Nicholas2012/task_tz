package randomuser

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"time"
)

const (
	defaultTimeout = 10 * time.Second
	baseURL        = "https://randomuser.me/api/"
)

type Client struct {
	http  *http.Client
	debug bool
}

func New() *Client {
	return &Client{
		http: &http.Client{
			Timeout: defaultTimeout,
		},
	}
}

func (c *Client) SetHTTPClient(client *http.Client) {
	c.http = client
}

func (c *Client) SetDebug(debug bool) {
	c.debug = debug
}

func (c *Client) Get() (*Response, error) {
	req, err := http.NewRequest(http.MethodGet, baseURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var reader io.Reader = resp.Body

	if c.debug {
		// for debug purposes
		reader = io.TeeReader(resp.Body, os.Stdout)
	}

	var r Response
	if err := json.NewDecoder(reader).Decode(&r); err != nil {
		return nil, err
	}

	return &r, nil
}
