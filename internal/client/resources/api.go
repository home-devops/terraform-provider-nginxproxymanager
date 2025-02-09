// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resources

type Api struct {
	Status  string  `json:"status"`
	Version Version `json:"version"`
}
