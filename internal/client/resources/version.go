// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resources

type Version struct {
	Major    int64 `json:"major"`
	Minor    int64 `json:"minor"`
	Revision int64 `json:"revision"`
}
