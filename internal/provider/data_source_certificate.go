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
	_ common.IDataSource                 = &certificateDataSource{}
	_ datasource.DataSourceWithConfigure = &certificateDataSource{}
)

func NewCertificateDataSource() datasource.DataSource {
	b := &common.DataSource{Name: "certificate"}
	d := &certificateDataSource{b, nil}
	b.IDataSource = d
	return d
}

type certificateDataSource struct {
	*common.DataSource
	client *client.Client
}

func (d *certificateDataSource) SchemaImpl(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "SSL Certificates --- Certificate data source",
		Attributes:  attributes.Certificate,
	}
}

func (d *certificateDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
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

func (d *certificateDataSource) ReadImpl(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {

	var data models.Certificate

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	certificate, err := d.client.GetCertificateCustom(ctx, data.ID.ValueInt64Pointer())
	if err != nil {
		resp.Diagnostics.AddError("Error reading certificate", "Could not read certificate, unexpected error: "+err.Error())
		return
	}
	if certificate == nil {
		resp.Diagnostics.AddError(
			"Error reading certificate",
			fmt.Sprintf("No certificate found with ID: %d", data.ID))
		return
	}

	resp.Diagnostics.Append(data.Load(ctx, certificate)...)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
}
