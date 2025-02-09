// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"terraform-provider-nginxproxymanager/internal/client/resources"
)

const usersUri = "%s/api/users?expand=permissions"
const userUri = "%s/api/users/%s?expand=permissions"

func (c *Client) GetUsers(ctx context.Context) (*resources.UserCollection, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf(usersUri, c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	ar := resources.UserCollection{}
	err = json.Unmarshal(body, &ar)
	if err != nil {
		return nil, err
	}

	return &ar, nil
}

func (c *Client) GetUser(ctx context.Context, id *int64) (*resources.User, error) {
	return c.getUser(ctx, strconv.FormatInt(*id, 10))
}

func (c *Client) GetMe(ctx context.Context) (*resources.User, error) {
	return c.getUser(ctx, "me")
}

func (c *Client) getUser(ctx context.Context, resource string) (*resources.User, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf(userUri, c.HostURL, resource), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	ar := resources.User{}
	err = json.Unmarshal(body, &ar)
	if err != nil {
		return nil, err
	}

	if ar.ID == 0 {
		return nil, nil
	}

	return &ar, nil
}
