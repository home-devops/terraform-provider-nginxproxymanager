// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package common

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var (
	_ datasource.DataSource = &DataSource{}
)

type IDataSource interface {
	SchemaImpl(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse)
	ReadImpl(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse)
}

type DataSource struct {
	IDataSource

	Name string
}

func (d *DataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = fmt.Sprintf("%s_%s", req.ProviderTypeName, d.Name)
}

func (d *DataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	d.SchemaImpl(ctx, req, resp)
}

func (d *DataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.ReadImpl(ctx, req, resp)
}
