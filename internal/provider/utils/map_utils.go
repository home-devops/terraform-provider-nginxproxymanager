// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package utils

import (
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func MergeMaps[K comparable, V any](maps ...map[K]V) map[K]V {
	result := make(map[K]V)
	for _, m := range maps {
		for k, v := range m {
			result[k] = v
		}
	}
	return result
}

func ConvertListToStringSlice(list types.List) (result []string, diags diag.Diagnostics) {
	for _, item := range list.Elements() {
		strVal, ok := item.(types.String)
		if !ok {
			diags.AddError("list contains non-string elements", strVal.String())
			return nil, diags
		}
		result = append(result, strVal.ValueString())
	}

	return result, nil
}
