// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resources

type AccessList struct {
	resourceWithOwner
	Name           string                     `json:"name"`
	Items          AccessListAuthCollection   `json:"items"`
	Clients        AccessListClientCollection `json:"clients"`
	PassAuth       bool                       `json:"pass_auth"`
	SatisfyAny     bool                       `json:"satisfy_any"`
	ProxyHostCount int64                      `json:"proxy_host_count"`
}

type AccessListCollection []AccessList
