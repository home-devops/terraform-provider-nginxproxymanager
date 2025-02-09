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
	_ common.IDataSource                 = &DeadHostsDataSource{}
	_ datasource.DataSourceWithConfigure = &DeadHostsDataSource{}
)

func NewDeadHostsDataSource() datasource.DataSource {
	b := &common.DataSource{Name: "dead_hosts"}
	d := &DeadHostsDataSource{b, nil}
	b.IDataSource = d
	return d
}

type DeadHostsDataSource struct {
	*common.DataSource
	client *client.Client
}

func (d *DeadHostsDataSource) SchemaImpl(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Hosts --- Dead Hosts data source",
		Attributes:  attributes.DeadHosts,
	}
}

func (d *DeadHostsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*client.Client)
}

func (d *DeadHostsDataSource) ReadImpl(ctx context.Context, _ datasource.ReadRequest, resp *datasource.ReadResponse) {
	deadHosts, err := d.client.GetDeadHosts(ctx)
	if err != nil {
		resp.Diagnostics.AddError("Error reading dead hosts", "Could not read dead hosts, unexpected error: "+err.Error())
		return
	}

	var data models.DeadHosts
	resp.Diagnostics.Append(data.Load(ctx, deadHosts)...)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
}
