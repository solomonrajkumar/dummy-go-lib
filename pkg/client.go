package dummylib

import (
	"net/http"
	"time"
)

const (
	BaseURLV1 = "http://localhost:8080/v1/organisation/accounts"
)

type Client struct {
	BaseURL    string
	HTTPClient *http.Client
}

func NewClient() *Client {
	return &Client{
		BaseURL: BaseURLV1,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}