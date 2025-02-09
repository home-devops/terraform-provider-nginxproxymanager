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
	_ common.IDataSource                 = &ProxyHostsDataSource{}
	_ datasource.DataSourceWithConfigure = &ProxyHostsDataSource{}
)

func NewProxyHostsDataSource() datasource.DataSource {
	b := &common.DataSource{Name: "proxy_hosts"}
	d := &ProxyHostsDataSource{b, nil}
	b.IDataSource = d
	return d
}

type ProxyHostsDataSource struct {
	*common.DataSource
	client *client.Client
}

func (d *ProxyHostsDataSource) SchemaImpl(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Hosts --- Proxy Hosts data source",
		Attributes:  attributes.ProxyHosts,
	}
}

func (d *ProxyHostsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*client.Client)
}

func (d *ProxyHostsDataSource) ReadImpl(ctx context.Context, _ datasource.ReadRequest, resp *datasource.ReadResponse) {
	proxyHosts, err := d.client.GetProxyHosts(ctx)
	if err != nil {
		resp.Diagnostics.AddError("Error reading proxy hosts", "Could not read proxy hosts, unexpected error: "+err.Error())
		return
	}

	var data models.ProxyHosts
	resp.Diagnostics.Append(data.Load(ctx, proxyHosts)...)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
}
