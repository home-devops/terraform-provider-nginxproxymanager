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
	_ common.IDataSource                 = &AccessListDataSource{}
	_ datasource.DataSourceWithConfigure = &AccessListDataSource{}
)

func NewAccessListDataSource() datasource.DataSource {
	b := &common.DataSource{Name: "access_list"}
	d := &AccessListDataSource{b, nil}
	b.IDataSource = d
	return d
}

type AccessListDataSource struct {
	*common.DataSource
	client *client.Client
}

func (d *AccessListDataSource) SchemaImpl(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Access Lists --- Access List data source",
		Attributes:  attributes.AccessList,
	}
}

func (d *AccessListDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*client.Client)
}

func (d *AccessListDataSource) ReadImpl(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data models.AccessList

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	accessList, err := d.client.GetAccessList(ctx, data.ID.ValueInt64Pointer())
	if err != nil {
		resp.Diagnostics.AddError("Error reading access list", "Could not read access list, unexpected error: "+err.Error())
		return
	}
	if accessList == nil {
		resp.Diagnostics.AddError(
			"Error reading access list",
			fmt.Sprintf("No access list found with ID: %d", data.ID))
		return
	}

	resp.Diagnostics.Append(data.Load(ctx, accessList)...)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
}
