# Create a letsencrypt certificate with DNS provider
resource "npm_certificate_letsencrypt" "example1" {
  domain_names             = ["example.com"]
  dns_challenge            = true
  letsencrypt_email        = "admin@example.com"
  dns_provider             = "cloudflare"
  dns_provider_credentials = "# Cloudflare API token\ndns_cloudflare_api_token=0123456789abcdef0123456789abcdef01234567"
  letsencrypt_agree        = true
}

# Create a letsencrypt certificate
resource "npm_certificate_letsencrypt" "example2" {
  domain_names      = ["example.com"]
  letsencrypt_email = "admin@example.com"
  letsencrypt_agree = true
}
