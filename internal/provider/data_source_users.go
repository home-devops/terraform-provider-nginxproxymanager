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
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ common.IDataSource                 = &UsersDataSource{}
	_ datasource.DataSourceWithConfigure = &UsersDataSource{}
)

func NewUsersDataSource() datasource.DataSource {
	b := &common.DataSource{Name: "users"}
	d := &UsersDataSource{b, nil}
	b.IDataSource = d
	return d
}

type UsersDataSource struct {
	*common.DataSource
	client *client.Client
}

func (d *UsersDataSource) SchemaImpl(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Users --- Users data source",
		Attributes:  attributes.Users,
	}
}

func (d *UsersDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
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

func (d *UsersDataSource) ReadImpl(ctx context.Context, _ datasource.ReadRequest, resp *datasource.ReadResponse) {
	users, err := d.client.GetUsers(ctx)
	if err != nil {
		resp.Diagnostics.AddError("Error reading user", "Could not read user, unexpected error: "+err.Error())
		return
	}

	var data models.Users
	resp.Diagnostics.Append(data.Load(ctx, users)...)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
}
