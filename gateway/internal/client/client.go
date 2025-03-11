package client

import (
	"net/http"
	"time"

	"github.com/go-jedi/gateway/config"
	"github.com/go-jedi/gateway/internal/client/posts"
)

const defaultTimeoutReq = 15 // second

// Client represents an HTTP client.
type Client struct {
	Posts *posts.Client
}

// New creates a new instance of HTTP client with a specified timeout.
func New(cfg config.ClientConfig) (client *Client, err error) {
	c := &Client{}

	httpClient := &http.Client{
		Timeout: time.Duration(cfg.TimeoutReq) * time.Second,
	}

	if cfg.TimeoutReq <= 0 {
		httpClient.Timeout = time.Duration(defaultTimeoutReq) * time.Second
	}

	c.Posts, err = posts.New(cfg, httpClient)
	if err != nil {
		return nil, err
	}

	return c, nil
}
