package models

import (
	"net/http"
	"time"
)

const (
	BaseURLV1 = "https://apidev.digitalxc.com/pb-dev/dxc/api"
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