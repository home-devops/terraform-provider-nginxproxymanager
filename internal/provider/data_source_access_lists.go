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
	_ common.IDataSource                 = &AccessListsDataSource{}
	_ datasource.DataSourceWithConfigure = &AccessListsDataSource{}
)

func NewAccessListsDataSource() datasource.DataSource {
	b := &common.DataSource{Name: "access_lists"}
	d := &AccessListsDataSource{b, nil}
	b.IDataSource = d
	return d
}

type AccessListsDataSource struct {
	*common.DataSource
	client *client.Client
}

func (d *AccessListsDataSource) SchemaImpl(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Access Lists --- Access Lists data source",
		Attributes:  attributes.AccessLists,
	}
}

func (d *AccessListsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*client.Client)
}

func (d *AccessListsDataSource) ReadImpl(ctx context.Context, _ datasource.ReadRequest, resp *datasource.ReadResponse) {
	accessLists, err := d.client.GetAccessLists(ctx)
	if err != nil {
		resp.Diagnostics.AddError("Error reading access lists", "Could not read access lists, unexpected error: "+err.Error())
		return
	}

	var data models.AccessLists
	resp.Diagnostics.Append(data.Load(ctx, accessLists)...)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
}
