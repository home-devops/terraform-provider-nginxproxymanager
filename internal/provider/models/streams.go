// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package models

import (
	"context"
	"terraform-provider-nginxproxymanager/internal/client/resources"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

type Streams struct {
	Streams []Stream `tfsdk:"streams"`
}

func (m *Streams) Load(ctx context.Context, resource *resources.StreamCollection) diag.Diagnostics {
	diags := diag.Diagnostics{}
	m.Streams = make([]Stream, len(*resource))
	for i, stream := range *resource {
		diags.Append(m.Streams[i].Load(ctx, &stream)...)
	}

	return diags
}
