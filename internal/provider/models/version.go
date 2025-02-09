// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package models

import (
	"context"
	"fmt"
	"terraform-provider-nginxproxymanager/internal/client/resources"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type Version struct {
	Major    types.Int64  `tfsdk:"major"`
	Minor    types.Int64  `tfsdk:"minor"`
	Revision types.Int64  `tfsdk:"revision"`
	Version  types.String `tfsdk:"version"`
}

func (m *Version) Load(_ context.Context, resource *resources.Api) diag.Diagnostics {
	diags := diag.Diagnostics{}

	m.Major = types.Int64Value(resource.Version.Major)
	m.Minor = types.Int64Value(resource.Version.Minor)
	m.Revision = types.Int64Value(resource.Version.Revision)
	m.Version = types.StringValue(fmt.Sprintf("%d.%d.%d", resource.Version.Major, resource.Version.Minor, resource.Version.Revision))

	return diags
}
