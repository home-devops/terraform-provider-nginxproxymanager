terraform {
  required_providers {
    npm = {
      source = "registry.opentofu.org/home-devops/nginxproxymanager"
    }
  }
}

provider "npm" {
  url      = "http://localhost:81"
  username = "terraform@tf.local"
  password = "password"
}

data "npm_redirection_hosts" "example" {}
data "npm_proxy_hosts" "example" {}


resource "npm_redirection_host" "test1" {
  domain_names        = ["example3.com"]
  forward_scheme      = "http"
  forward_domain_name = "example.com"
  forward_http_code   = 301
  hsts_enabled        = true
  certificate_id      = 1
  ssl_forced          = true
}

resource "npm_redirection_host" "test2" {
  domain_names        = ["example.com"]
  forward_scheme      = "http"
  forward_domain_name = "example2.com"
  forward_http_code   = 301
  certificate_id      = 1
}

resource "npm_proxy_host" "test1" {
  domain_names   = ["example1.com"]
  forward_scheme = "http"
  forward_host   = "example.com"
  forward_port   = 80
}

data "npm_proxy_host" "test2" {
  id = 2
}

# data "npm_certificates" "test" {}
data "npm_certificate" "test" {
  id = 1
}

resource "npm_certificate_letsencrypt" "test" {
  domain_names             = ["krivopishin.by"]
  dns_challenge            = true
  letsencrypt_email        = "s.krivopishin@gmail.com"
  dns_provider             = "cloudflare"
  dns_provider_credentials = "dns_cloudflare_api_token=yClPbfNui0Vu7RiI_zLg5B6oU8-9B7k00MnM2LKH"
  letsencrypt_agree        = true
}

resource "npm_certificate_letsencrypt" "example2" {
  domain_names      = ["example.com"]
  letsencrypt_email = "admin@example.com"
  letsencrypt_agree = true
}

data "npm_access_lists" "test" {}
data "npm_dead_hosts" "test" {}
data "npm_streams" "test" {}
data "npm_users" "test" {}
data "npm_access_list" "test" {
  id = 1
}
data "npm_dead_host" "test" {
  id = 1
}
data "npm_stream" "test" {
  id = 1
}
data "npm_user" "test" {
  id = 1
}
data "npm_user_me" "test" {}
data "npm_version" "test" {}

output "test" {
  value = jsondecode(data.npm_certificate.test.meta.certificate)
}
