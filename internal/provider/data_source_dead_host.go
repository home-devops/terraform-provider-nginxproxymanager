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
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ common.IDataSource                 = &DeadHostDataSource{}
	_ datasource.DataSourceWithConfigure = &DeadHostDataSource{}
)

func NewDeadHostDataSource() datasource.DataSource {
	b := &common.DataSource{Name: "dead_host"}
	d := &DeadHostDataSource{b, nil}
	b.IDataSource = d
	return d
}

type DeadHostDataSource struct {
	*common.DataSource
	client *client.Client
}

func (d *DeadHostDataSource) SchemaImpl(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Hosts --- Fetches a dead host by ID.",
		Attributes:  attributes.DeadHost,
	}
}

func (d *DeadHostDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	var ok bool
	d.client, ok = req.ProviderData.(*client.Client)
	if !ok {
		tflog.Error(ctx, "ProviderData is not of type *client.Client")
		return
	}
}

func (d *DeadHostDataSource) ReadImpl(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data models.DeadHost

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	deadHost, err := d.client.GetDeadHost(ctx, data.ID.ValueInt64Pointer())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading dead host",
			"Could not read dead host, unexpected error: "+err.Error())
		return
	}
	if deadHost == nil {
		resp.Diagnostics.AddError(
			"Error reading dead host",
			fmt.Sprintf("No dead host found with ID: %d", data.ID))
		return
	}

	resp.Diagnostics.Append(data.Load(ctx, deadHost)...)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
}
