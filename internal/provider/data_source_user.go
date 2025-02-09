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
	_ common.IDataSource                 = &UserDataSource{}
	_ datasource.DataSourceWithConfigure = &UserDataSource{}
)

func NewUserDataSource() datasource.DataSource {
	b := &common.DataSource{Name: "user"}
	d := &UserDataSource{b, nil}
	b.IDataSource = d
	return d
}

type UserDataSource struct {
	*common.DataSource
	client *client.Client
}

func (d *UserDataSource) SchemaImpl(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Users --- User data source",
		Attributes:  attributes.User,
	}
}

func (d *UserDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
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

func (d *UserDataSource) ReadImpl(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data models.User

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	user, err := d.client.GetUser(ctx, data.ID.ValueInt64Pointer())
	if err != nil {
		resp.Diagnostics.AddError("Error reading user", "Could not read user, unexpected error: "+err.Error())
		return
	}
	if user == nil {
		resp.Diagnostics.AddError(
			"Error reading user",
			fmt.Sprintf("No user found with ID: %d", data.ID))
		return
	}

	resp.Diagnostics.Append(data.Load(ctx, user)...)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
}
