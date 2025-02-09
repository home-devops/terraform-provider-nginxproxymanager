// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package inputs

type CertificateUpload struct {
	CertificateId  int64  `json:"certificate_id"`
	Certificate    string `json:"certificate"`
	CertificateKey string `json:"certificate_key"`
}
