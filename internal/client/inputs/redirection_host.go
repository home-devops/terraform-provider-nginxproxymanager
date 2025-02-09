// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package inputs

type RedirectionHost struct {
	DomainNames       []string          `json:"domain_names"`
	ForwardScheme     string            `json:"forward_scheme"`
	ForwardDomainName string            `json:"forward_domain_name"`
	ForwardHTTPCode   int64             `json:"forward_http_code"`
	CertificateID     string            `json:"certificate_id"`
	SSLForced         bool              `json:"ssl_forced"`
	HSTSEnabled       bool              `json:"hsts_enabled"`
	HSTSSubdomains    bool              `json:"hsts_subdomains"`
	HTTP2Support      bool              `json:"http2_support"`
	PreservePath      bool              `json:"preserve_path"`
	BlockExploits     bool              `json:"block_exploits"`
	AdvancedConfig    string            `json:"advanced_config"`
	Meta              map[string]string `json:"meta"`
}
