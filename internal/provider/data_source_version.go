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
	_ common.IDataSource                 = &VersionDataSource{}
	_ datasource.DataSourceWithConfigure = &VersionDataSource{}
)

func NewVersionDataSource() datasource.DataSource {
	b := &common.DataSource{Name: "version"}
	d := &VersionDataSource{b, nil}
	b.IDataSource = d
	return d
}

type VersionDataSource struct {
	*common.DataSource
	client *client.Client
}

func (d *VersionDataSource) SchemaImpl(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Meta --- Version data source",
		Attributes:  attributes.Version,
	}
}

func (d *VersionDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*client.Client)
}

func (d *VersionDataSource) ReadImpl(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data models.Version

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	api, err := d.client.GetApi(ctx)
	if err != nil {
		resp.Diagnostics.AddError("Error reading version", "Could not read version, unexpected error: "+err.Error())
		return
	}
	if api == nil {
		resp.Diagnostics.AddError("Error reading version", "No version found")
		return
	}

	resp.Diagnostics.Append(data.Load(ctx, api)...)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
}
