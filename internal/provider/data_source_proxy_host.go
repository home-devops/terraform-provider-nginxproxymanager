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
	_ common.IDataSource                 = &ProxyHostDataSource{}
	_ datasource.DataSourceWithConfigure = &ProxyHostDataSource{}
)

func NewProxyHostDataSource() datasource.DataSource {
	b := &common.DataSource{Name: "proxy_host"}
	d := &ProxyHostDataSource{b, nil}
	b.IDataSource = d
	return d
}

type ProxyHostDataSource struct {
	*common.DataSource
	client *client.Client
}

func (d *ProxyHostDataSource) SchemaImpl(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Hosts --- Fetches a proxy host by ID.",
		Attributes:  attributes.ProxyHost,
	}
}

func (d *ProxyHostDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*client.Client)
}

func (d *ProxyHostDataSource) ReadImpl(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data models.ProxyHost

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	proxyHost, err := d.client.GetProxyHost(ctx, data.ID.ValueInt64Pointer())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading proxy host",
			"Could not read proxy host, unexpected error: "+err.Error())
		return
	}
	if proxyHost == nil {
		resp.Diagnostics.AddError(
			"Error reading proxy host",
			fmt.Sprintf("No proxy host found with ID: %d", data.ID))
		return
	}

	resp.Diagnostics.Append(data.Load(ctx, proxyHost)...)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
}
