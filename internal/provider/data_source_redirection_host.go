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
	_ common.IDataSource                 = &RedirectionHostDataSource{}
	_ datasource.DataSourceWithConfigure = &RedirectionHostDataSource{}
)

func NewRedirectionHostDataSource() datasource.DataSource {
	b := &common.DataSource{Name: "redirection_host"}
	d := &RedirectionHostDataSource{b, nil}
	b.IDataSource = d
	return d
}

type RedirectionHostDataSource struct {
	*common.DataSource
	client *client.Client
}

func (d *RedirectionHostDataSource) SchemaImpl(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Hosts --- Fetches a redirection host by ID.",
		Attributes:  attributes.RedirectionHost,
	}
}

func (d *RedirectionHostDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*client.Client)
}

func (d *RedirectionHostDataSource) ReadImpl(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data models.RedirectionHost

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	redirectionHost, err := d.client.GetRedirectionHost(ctx, data.ID.ValueInt64Pointer())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading redirection host",
			"Could not read redirection host, unexpected error: "+err.Error())
		return
	}
	if redirectionHost == nil {
		resp.Diagnostics.AddError(
			"Error reading redirection host",
			fmt.Sprintf("No redirection host found with ID: %d", data.ID))
		return
	}

	resp.Diagnostics.Append(data.Load(ctx, redirectionHost)...)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
}
