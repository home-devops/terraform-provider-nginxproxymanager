// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resources

import (
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var RedirectionHost = map[string]schema.Attribute{
	"id": schema.Int64Attribute{
		Description: "The ID of the redirection host.",
		Computed:    true,
		PlanModifiers: []planmodifier.Int64{
			int64planmodifier.UseStateForUnknown(),
		},
	},
	"created_on": schema.StringAttribute{
		Description: "The date and time the redirection host was created.",
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	},
	"modified_on": schema.StringAttribute{
		Description: "The date and time the redirection host was last modified.",
		Computed:    true,
	},
	"owner_user_id": schema.Int64Attribute{
		Description: "The ID of the user that owns the redirection host.",
		Computed:    true,
		PlanModifiers: []planmodifier.Int64{
			int64planmodifier.UseStateForUnknown(),
		},
	},
	"domain_names": schema.ListAttribute{
		Description: "Domain Names separated by a comma.",
		Required:    true,
		ElementType: types.StringType,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
			listvalidator.SizeAtMost(100),
			listvalidator.UniqueValues(),
			listvalidator.ValueStringsAre(stringvalidator.RegexMatches(regexp.MustCompile("^[^&| @!#%^();:/\\\\}{=+?<>,~`'\"]+$"), "Incorrect value.")),
		},
	},
	"forward_scheme": schema.StringAttribute{
		Description: "The scheme used to forward requests to the redirection host. Can be either `http` or `https` or `auto`.",
		Required:    true,
		Validators: []validator.String{
			stringvalidator.OneOf("auto", "http", "https"),
		},
	},
	"forward_http_code": schema.Int64Attribute{
		Description: "Redirect HTTP Status Code.",
		Required:    true,
		Validators: []validator.Int64{
			int64validator.Between(300, 308),
		},
	},
	"forward_domain_name": schema.StringAttribute{
		Description: "Domain Name.",
		Required:    true,
		Validators: []validator.String{
			stringvalidator.RegexMatches(regexp.MustCompile(`^(?:[^.*]+\.?)+[^.]$`), "Incorrect value."),
		},
	},
	"preserve_path": schema.BoolAttribute{
		Description: "Should the path be preserved.",
		Computed:    true,
		Optional:    true,
		Default:     booldefault.StaticBool(false),
	},
	"certificate_id": schema.Int64Attribute{
		Description: "Certificate ID.",
		Computed:    true,
		Optional:    true,
	},
	"certificate_new": schema.BoolAttribute{
		Description: "Generate certificate using HTTP.",
		Computed:    true,
		Optional:    true,
		Default:     booldefault.StaticBool(false),
	},
	"ssl_forced": schema.BoolAttribute{
		Description: "Whether SSL is forced for the proxy host.",
		Computed:    true,
		Optional:    true,
		Default:     booldefault.StaticBool(false),
	},
	"hsts_enabled": schema.BoolAttribute{
		Description: "Whether HSTS is enabled for the proxy host.",
		Computed:    true,
		Optional:    true,
		Default:     booldefault.StaticBool(false),
	},
	"hsts_subdomains": schema.BoolAttribute{
		Description: "Whether HSTS is enabled for subdomains of the proxy host.",
		Computed:    true,
		Optional:    true,
		Default:     booldefault.StaticBool(false),
	},
	"http2_support": schema.BoolAttribute{
		Description: "Whether HTTP/2 is supported for the proxy host.",
		Computed:    true,
		Optional:    true,
		Default:     booldefault.StaticBool(false),
	},
	"block_exploits": schema.BoolAttribute{
		Description: "Should we block common exploits.",
		Computed:    true,
		Optional:    true,
		Default:     booldefault.StaticBool(false),
	},
	"advanced_config": schema.StringAttribute{
		Description: "The advanced configuration used by the proxy host.",
		Computed:    true,
		Optional:    true,
		Default:     stringdefault.StaticString(""),
	},
	"enabled": schema.BoolAttribute{
		Description: "Whether the redirection host is enabled.",
		Computed:    true,
		PlanModifiers: []planmodifier.Bool{
			boolplanmodifier.UseStateForUnknown(),
		},
	},
	"meta": schema.MapAttribute{
		Description: "The meta data associated with the proxy host.",
		ElementType: types.StringType,
		Computed:    true,
	},
}
