---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "npm_redirection_host Resource - npm"
subcategory: ""
description: |-
  Hosts --- Manage a redirection host.
---

# npm_redirection_host (Resource)

Hosts --- Manage a redirection host.

## Example Usage

```terraform
# Create a redirection host
resource "npm_redirection_host" "example" {
  domain_names = ["example.com"]

  forward_scheme      = "http"
  forward_domain_name = "example.com"
  forward_http_code   = 301

  preserve_path = true

  certificate_id = 0 # No certificate

  hsts_enabled    = false
  hsts_subdomains = false
  http2_support   = false
  ssl_forced      = false

  block_exploits = true

  advanced_config = ""
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `domain_names` (List of String) Domain Names separated by a comma.
- `forward_domain_name` (String) Domain Name.
- `forward_http_code` (Number) Redirect HTTP Status Code.
- `forward_scheme` (String) The scheme used to forward requests to the redirection host. Can be either `http` or `https` or `auto`.

### Optional

- `advanced_config` (String) The advanced configuration used by the proxy host.
- `block_exploits` (Boolean) Should we block common exploits.
- `certificate_id` (Number) Certificate ID.
- `certificate_new` (Boolean) Generate certificate using HTTP.
- `hsts_enabled` (Boolean) Whether HSTS is enabled for the proxy host.
- `hsts_subdomains` (Boolean) Whether HSTS is enabled for subdomains of the proxy host.
- `http2_support` (Boolean) Whether HTTP/2 is supported for the proxy host.
- `preserve_path` (Boolean) Should the path be preserved.
- `ssl_forced` (Boolean) Whether SSL is forced for the proxy host.

### Read-Only

- `created_on` (String) The date and time the redirection host was created.
- `enabled` (Boolean) Whether the redirection host is enabled.
- `id` (Number) The ID of the redirection host.
- `meta` (Map of String) The meta data associated with the proxy host.
- `modified_on` (String) The date and time the redirection host was last modified.
- `owner_user_id` (Number) The ID of the user that owns the redirection host.

## Import

Import is supported using the following syntax:

```shell
# Redirection hosts can be imported by specifying the numeric identifier of the redirection host.
terraform import npm_redirection_hosts.example 1
```
