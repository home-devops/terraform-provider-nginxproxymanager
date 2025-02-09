// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package inputs

type ProxyHostLocation struct {
	Path           string `json:"path"`
	ForwardScheme  string `json:"forward_scheme"`
	ForwardHost    string `json:"forward_host"`
	ForwardPort    uint16 `json:"forward_port"`
	AdvancedConfig string `json:"advanced_config"`
}

type ProxyHostLocationCollection []ProxyHostLocation
