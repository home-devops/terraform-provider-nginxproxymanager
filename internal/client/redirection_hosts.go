// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"terraform-provider-nginxproxymanager/internal/client/inputs"
	"terraform-provider-nginxproxymanager/internal/client/resources"

	"github.com/go-http-utils/headers"
)

const redirectionHostsUri = "%s/api/nginx/redirection-hosts"
const redirectionHostUri = "%s/api/nginx/redirection-hosts/%d"

func (c *Client) GetRedirectionHosts(ctx context.Context) (*resources.RedirectionHostCollection, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf(redirectionHostsUri, c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	ar := resources.RedirectionHostCollection{}
	err = json.Unmarshal(body, &ar)
	if err != nil {
		return nil, err
	}

	return &ar, nil
}

func (c *Client) GetRedirectionHost(ctx context.Context, id *int64) (*resources.RedirectionHost, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf(redirectionHostUri, c.HostURL, *id), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	ar := resources.RedirectionHost{}
	err = json.Unmarshal(body, &ar)
	if err != nil {
		return nil, err
	}

	if ar.ID == 0 {
		return nil, nil
	}

	return &ar, nil
}

func (c *Client) CreateRedirectionHost(ctx context.Context, redirectionHost *inputs.RedirectionHost) (*resources.RedirectionHost, error) {
	rb, err := json.Marshal(redirectionHost)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", fmt.Sprintf(redirectionHostsUri, c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	req.Header.Set(headers.ContentType, "application/json")

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	ar := resources.RedirectionHost{}
	err = json.Unmarshal(body, &ar)
	if err != nil {
		return nil, err
	}

	return &ar, nil
}

func (c *Client) UpdateRedirectionHost(ctx context.Context, id *int64, redirectionHost *inputs.RedirectionHost) (*resources.RedirectionHost, error) {
	rb, err := json.Marshal(redirectionHost)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "PUT", fmt.Sprintf(redirectionHostUri, c.HostURL, *id), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	req.Header.Set(headers.ContentType, "application/json")

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	ar := resources.RedirectionHost{}
	err = json.Unmarshal(body, &ar)
	if err != nil {
		return nil, err
	}

	return &ar, nil
}

func (c *Client) DeleteRedirectionHost(ctx context.Context, id *int64) error {
	req, err := http.NewRequestWithContext(ctx, "DELETE", fmt.Sprintf(redirectionHostUri, c.HostURL, *id), nil)
	if err != nil {
		return err
	}

	_, err = c.doRequest(req)
	if err != nil {
		return err
	}

	return nil
}
