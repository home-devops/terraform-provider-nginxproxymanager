// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resources

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var CertificateLetsEncrypt = map[string]schema.Attribute{
	"id": schema.Int64Attribute{
		Description: "The ID of the certificate.",
		Computed:    true,
		PlanModifiers: []planmodifier.Int64{
			int64planmodifier.UseStateForUnknown(),
		},
	},
	"name": schema.StringAttribute{
		Description: "The name of the certificate.",
		Computed:    true,
	},
	"dns_challenge": schema.BoolAttribute{
		Description: "Whether DNS challange should be used.",
		Optional:    true,
		PlanModifiers: []planmodifier.Bool{
			boolplanmodifier.RequiresReplace(),
		},
	},
	"letsencrypt_email": schema.StringAttribute{
		Description: "Letsencrypt user email to use.",
		Required:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.RequiresReplace(),
		},
	},
	"dns_provider": schema.StringAttribute{
		Description: "DNS provider name to use.",
		Optional:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.RequiresReplace(),
		},
	},
	"dns_provider_credentials": schema.StringAttribute{
		Description: "DNS provider credentials to use.",
		Optional:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.RequiresReplace(),
		},
	},
	"letsencrypt_agree": schema.BoolAttribute{
		Description: "Whether agree letsencrypt terms of service.",
		Required:    true,
		PlanModifiers: []planmodifier.Bool{
			boolplanmodifier.RequiresReplace(),
		},
	},
	"domain_names": schema.ListAttribute{
		Description: "The domain names associated with the certificate.",
		Required:    true,
		ElementType: types.StringType,
		PlanModifiers: []planmodifier.List{
			listplanmodifier.RequiresReplace(),
		},
	},
	"expires_on": schema.StringAttribute{
		Description: "The date and time the certificate expires.",
		Computed:    true,
	},
	"created_on": schema.StringAttribute{
		Description: "The date and time the certificate was created.",
		Computed:    true,
	},
	"modified_on": schema.StringAttribute{
		Description: "The date and time the certificate was last modified.",
		Computed:    true,
	},
}
