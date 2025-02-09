// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"terraform-provider-nginxproxymanager/internal/client"
	attributes "terraform-provider-nginxproxymanager/internal/provider/attributes/datasource"
	"terraform-provider-nginxproxymanager/internal/provider/common"
	"terraform-provider-nginxproxymanager/internal/provider/models"
)

var _ datasource.DataSource = &RedirectionHostsDataSource{}

func NewRedirectionHostsDataSource() datasource.DataSource {
	b := &common.DataSource{Name: "redirection_hosts"}
	d := &RedirectionHostsDataSource{b, nil}
	b.IDataSource = d
	return d
}

type RedirectionHostsDataSource struct {
	*common.DataSource
	client *client.Client
}

func (d *RedirectionHostsDataSource) SchemaImpl(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Hosts --- Redirection Hosts data source",
		Attributes:  attributes.RedirectionHosts,
	}
}

func (d *RedirectionHostsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*client.Client)
}

func (d *RedirectionHostsDataSource) ReadImpl(ctx context.Context, _ datasource.ReadRequest, resp *datasource.ReadResponse) {
	redirectionHosts, err := d.client.GetRedirectionHosts(ctx)
	if err != nil {
		resp.Diagnostics.AddError("Error reading redirection hosts", "Could not read redirection hosts, unexpected error: "+err.Error())
		return
	}

	var data models.RedirectionHosts
	resp.Diagnostics.Append(data.Load(ctx, redirectionHosts)...)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
}
