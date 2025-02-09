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

const streamsUri = "%s/api/nginx/streams"
const streamUri = "%s/api/nginx/streams/%d"

func (c *Client) GetStreams(ctx context.Context) (*resources.StreamCollection, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf(streamsUri, c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	ar := resources.StreamCollection{}
	err = json.Unmarshal(body, &ar)
	if err != nil {
		return nil, err
	}

	return &ar, nil
}

func (c *Client) GetStream(ctx context.Context, id *int64) (*resources.Stream, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf(streamUri, c.HostURL, *id), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	ar := resources.Stream{}
	err = json.Unmarshal(body, &ar)
	if err != nil {
		return nil, err
	}

	if ar.ID == 0 {
		return nil, nil
	}

	return &ar, nil
}
