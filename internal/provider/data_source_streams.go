// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"terraform-provider-nginxproxymanager/internal/provider/models"

	attributes "terraform-provider-nginxproxymanager/internal/provider/attributes/datasource"

	"terraform-provider-nginxproxymanager/internal/client"
	"terraform-provider-nginxproxymanager/internal/provider/common"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var (
	_ common.IDataSource                 = &StreamsDataSource{}
	_ datasource.DataSourceWithConfigure = &StreamsDataSource{}
)

func NewStreamsDataSource() datasource.DataSource {
	b := &common.DataSource{Name: "streams"}
	d := &StreamsDataSource{b, nil}
	b.IDataSource = d
	return d
}

type StreamsDataSource struct {
	*common.DataSource
	client *client.Client
}

func (d *StreamsDataSource) SchemaImpl(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Hosts --- Stream data source.",
		Attributes:  attributes.Streams,
	}
}

func (d *StreamsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*client.Client)
}

func (d *StreamsDataSource) ReadImpl(ctx context.Context, _ datasource.ReadRequest, resp *datasource.ReadResponse) {
	streams, err := d.client.GetStreams(ctx)
	if err != nil {
		resp.Diagnostics.AddError("Error reading streams", "Could not read streams, unexpected error: "+err.Error())
		return
	}

	var data models.Streams
	resp.Diagnostics.Append(data.Load(ctx, streams)...)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
}
