// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"terraform-provider-nginxproxymanager/internal/client/resources"
)

const deadHostsUri = "%s/api/nginx/dead-hosts"
const deadHostUri = "%s/api/nginx/dead-hosts/%d"

func (c *Client) GetDeadHosts(ctx context.Context) (*resources.DeadHostCollection, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf(deadHostsUri, c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	ar := resources.DeadHostCollection{}
	err = json.Unmarshal(body, &ar)
	if err != nil {
		return nil, err
	}

	return &ar, nil
}

func (c *Client) GetDeadHost(ctx context.Context, id *int64) (*resources.DeadHost, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf(deadHostUri, c.HostURL, *id), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	ar := resources.DeadHost{}
	err = json.Unmarshal(body, &ar)
	if err != nil {
		return nil, err
	}

	if ar.ID == 0 {
		return nil, nil
	}

	return &ar, nil
}
