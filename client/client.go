package client

import (
	"net/http"
)

type Client struct {
	httpclient *http.Client
	token     string //Optional
	base_url  string //Required
}

var clietnImpl *Client

type Option func(*Client)

func (c *Client) makeRequest(method, endpoint string, con *container.Container) (*http.Request, error) {
	var req *http.Request
	var err error
	if method == "POST" || method == "PUT" {
		req, err = http.NewRequest(method, endpoint, bytes.NewBuffer(con.Bytes()))
	} else {
		req, err = http.NewRequest(method, endpoint, nil)
	}
	if err != nil {
		return nil, err
	}

	return req, nil
}
