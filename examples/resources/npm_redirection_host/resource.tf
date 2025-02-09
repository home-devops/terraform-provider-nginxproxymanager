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
