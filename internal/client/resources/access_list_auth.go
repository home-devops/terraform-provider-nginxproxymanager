// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resources

type AccessListAuth struct {
	resource
	AccessListId int64  `json:"access_list_id"`
	Hint         string `json:"hint"`
	Username     string `json:"username"`
	Password     string `json:"password"`
}

type AccessListAuthCollection []AccessListAuth
