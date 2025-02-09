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

const accessListsUri = "%s/api/nginx/access-lists?expand=items,clients"
const accessListUri = "%s/api/nginx/access-lists/%d?expand=items,clients"

func (c *Client) GetAccessLists(ctx context.Context) (*resources.AccessListCollection, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf(accessListsUri, c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	ar := resources.AccessListCollection{}
	err = json.Unmarshal(body, &ar)
	if err != nil {
		return nil, err
	}

	return &ar, nil
}

func (c *Client) GetAccessList(ctx context.Context, id *int64) (*resources.AccessList, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf(accessListUri, c.HostURL, *id), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	ar := resources.AccessList{}
	err = json.Unmarshal(body, &ar)
	if err != nil {
		return nil, err
	}

	if ar.ID == 0 {
		return nil, nil
	}

	return &ar, nil
}
