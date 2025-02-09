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
	_ common.IDataSource                 = &certificatesDataSource{}
	_ datasource.DataSourceWithConfigure = &certificatesDataSource{}
)

func NewCertificatesDataSource() datasource.DataSource {
	b := &common.DataSource{Name: "certificates"}
	d := &certificatesDataSource{b, nil}
	b.IDataSource = d
	return d
}

type certificatesDataSource struct {
	*common.DataSource
	client *client.Client
}

func (d *certificatesDataSource) SchemaImpl(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "SSL Certificates --- Certificates data source",
		Attributes:  attributes.Certificates,
	}
}

func (d *certificatesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
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

func (d *certificatesDataSource) ReadImpl(ctx context.Context, _ datasource.ReadRequest, resp *datasource.ReadResponse) {
	certificates, err := d.client.GetCertificates(ctx)
	if err != nil {
		resp.Diagnostics.AddError("Error reading certificate", "Could not read certificate, unexpected error: "+err.Error())
		return
	}

	var data models.Certificates
	resp.Diagnostics.Append(data.Load(ctx, certificates)...)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
}
