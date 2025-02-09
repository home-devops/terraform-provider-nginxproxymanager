// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package common

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var (
	_ resource.Resource = &Resource{}
)

type IResource interface {
	SchemaImpl(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse)
	ReadImpl(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse)
	CreateImpl(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse)
	UpdateImpl(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse)
	DeleteImpl(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse)
}

type Resource struct {
	IResource

	Name string
}

func (r *Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = fmt.Sprintf("%s_%s", req.ProviderTypeName, r.Name)
}

func (r *Resource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	r.SchemaImpl(ctx, req, resp)
}

func (r *Resource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.ReadImpl(ctx, req, resp)

}

func (r *Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.CreateImpl(ctx, req, resp)
}

func (r *Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.UpdateImpl(ctx, req, resp)
}

func (r *Resource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	r.DeleteImpl(ctx, req, resp)
}
