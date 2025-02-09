// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"fmt"
	"terraform-provider-nginxproxymanager/internal/provider/models"

	attributes "terraform-provider-nginxproxymanager/internal/provider/attributes/datasource"

	"terraform-provider-nginxproxymanager/internal/client"
	"terraform-provider-nginxproxymanager/internal/provider/common"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var (
	_ common.IDataSource                 = &StreamDataSource{}
	_ datasource.DataSourceWithConfigure = &StreamDataSource{}
)

func NewStreamDataSource() datasource.DataSource {
	b := &common.DataSource{Name: "stream"}
	d := &StreamDataSource{b, nil}
	b.IDataSource = d
	return d
}

type StreamDataSource struct {
	*common.DataSource
	client *client.Client
}

func (d *StreamDataSource) SchemaImpl(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Hosts --- Fetches a stream by ID.",
		Attributes:  attributes.Stream,
	}
}

func (d *StreamDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*client.Client)
}

func (d *StreamDataSource) ReadImpl(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data models.Stream

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	stream, err := d.client.GetStream(ctx, data.ID.ValueInt64Pointer())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading stream",
			"Could not read stream, unexpected error: "+err.Error())
		return
	}
	if stream == nil {
		resp.Diagnostics.AddError(
			"Error reading stream",
			fmt.Sprintf("No stream found with ID: %d", data.ID))
		return
	}

	resp.Diagnostics.Append(data.Load(ctx, stream)...)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
}
