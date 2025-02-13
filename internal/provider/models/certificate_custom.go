// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package models

import (
	"context"
	"strings"
	"terraform-provider-nginxproxymanager/internal/client/inputs"
	"terraform-provider-nginxproxymanager/internal/client/resources"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CertificateCustom struct {
	ID         types.Int64  `tfsdk:"id"`
	CreatedOn  types.String `tfsdk:"created_on"`
	ModifiedOn types.String `tfsdk:"modified_on"`

	Name           types.String `tfsdk:"name"`
	DomainNames    types.List   `tfsdk:"domain_names"`
	ExpiresOn      types.String `tfsdk:"expires_on"`
	Certificate    types.String `tfsdk:"certificate"`
	CertificateKey types.String `tfsdk:"certificate_key"`
}

func (m *CertificateCustom) Load(ctx context.Context, resource *resources.Certificate) diag.Diagnostics {
	var diags diag.Diagnostics

	m.ID = types.Int64Value(resource.ID)
	m.CreatedOn = types.StringValue(resource.CreatedOn)
	m.ModifiedOn = types.StringValue(resource.ModifiedOn)
	m.Name = types.StringValue(resource.NiceName)
	m.ExpiresOn = types.StringValue(resource.ExpiresOn)
	m.Certificate = types.StringValue(strings.Trim(strings.ReplaceAll(resource.Meta.Map()["certificate"], "\\n", "\n"), "\""))
	m.CertificateKey = types.StringValue(strings.Trim(strings.ReplaceAll(resource.Meta.Map()["certificate_key"], "\\n", "\n"), "\""))

	m.DomainNames, diags = types.ListValueFrom(ctx, types.StringType, resource.DomainNames)

	if diags.HasError() {
		return diags
	}

	return diags
}

func (m *CertificateCustom) Save(_ context.Context, input *inputs.CertificateCustom) diag.Diagnostics {
	var diags diag.Diagnostics

	input.Name = m.Name.ValueString()
	input.Certificate = m.Certificate.ValueString()
	input.CertificateKey = m.CertificateKey.ValueString()

	return diags
}
