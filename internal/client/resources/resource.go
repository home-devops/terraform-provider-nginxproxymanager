// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resources

type resource struct {
	ID         int64  `json:"id"`
	CreatedOn  string `json:"created_on"`
	ModifiedOn string `json:"modified_on"`
	Meta       Meta   `json:"meta"`
}

type resourceWithOwner struct {
	resource
	OwnerUserID int64 `json:"owner_user_id"`
}
