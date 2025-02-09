// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package inputs

type CertificateLetsEncryptMeta struct {
	DnsChallenge           bool   `json:"dns_challenge"`
	DnsProvider            string `json:"dns_provider"`
	DnsProviderCredentials string `json:"dns_provider_credentials"`
	LetsEncryptAgree       bool   `json:"letsencrypt_agree"`
	LetsEncryptEmail       string `json:"letsencrypt_email"`
}

type CertificateLetsEncrypt struct {
	Provider    string                     `json:"provider"`
	DomainNames []string                   `json:"domain_names"`
	Meta        CertificateLetsEncryptMeta `json:"meta"`
}
