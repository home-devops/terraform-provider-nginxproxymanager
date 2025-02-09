// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package models

import (
	"context"
	"terraform-provider-nginxproxymanager/internal/client/resources"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

type Certificates struct {
	Certificates []Certificate `tfsdk:"certificates"`
}

func (m *Certificates) Load(ctx context.Context, resource *resources.CertificateCollection) diag.Diagnostics {
	diags := diag.Diagnostics{}
	m.Certificates = make([]Certificate, len(*resource))
	for i, certificate := range *resource {
		diags.Append(m.Certificates[i].Load(ctx, &certificate)...)
	}

	return diags
}
