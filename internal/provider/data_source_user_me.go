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
	_ common.IDataSource                 = &UserMeDataSource{}
	_ datasource.DataSourceWithConfigure = &UserMeDataSource{}
)

func NewUserMeDataSource() datasource.DataSource {
	b := &common.DataSource{Name: "user_me"}
	d := &UserMeDataSource{b, nil}
	b.IDataSource = d
	return d
}

type UserMeDataSource struct {
	*common.DataSource
	client *client.Client
}

func (d *UserMeDataSource) SchemaImpl(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Users --- User me data source",
		Attributes:  attributes.UserMe,
	}
}

func (d *UserMeDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*client.Client)
}

func (d *UserMeDataSource) ReadImpl(ctx context.Context, _ datasource.ReadRequest, resp *datasource.ReadResponse) {
	user, err := d.client.GetMe(ctx)
	if err != nil {
		resp.Diagnostics.AddError("Error reading user", "Could not read user, unexpected error: "+err.Error())
		return
	}

	var data models.User
	resp.Diagnostics.Append(data.Load(ctx, user)...)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
}
