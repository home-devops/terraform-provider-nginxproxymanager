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

func (c *Client) GetApi(ctx context.Context) (*resources.Api, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("%s/api", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	ar := resources.Api{}
	err = json.Unmarshal(body, &ar)
	if err != nil {
		return nil, err
	}

	return &ar, nil
}
