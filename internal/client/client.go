// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package client

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/go-http-utils/headers"
)

type Client struct {
	HostURL    string
	HTTPClient *http.Client
	Token      string
	UserAgent  string
}

func NewClient(url *string, username *string, password *string, version string) (*Client, error) {
	c := Client{
		HTTPClient: &http.Client{
			Timeout:   300 * time.Second,
			Transport: http.DefaultTransport,
		},
		UserAgent: fmt.Sprintf("terraform-provider-nginxproxymanager/%s", version),
	}

	if url != nil {
		c.HostURL = *url
	}

	if username == nil || password == nil {
		return &c, nil
	}

	ar, err := c.Authenticate(context.Background(), username, password)
	if err != nil {
		return nil, err
	}

	c.Token = ar.Token

	return &c, nil
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	token := c.Token

	req.Header.Set(headers.Authorization, fmt.Sprintf("Bearer %s", token))
	req.Header.Set(headers.UserAgent, c.UserAgent)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}
