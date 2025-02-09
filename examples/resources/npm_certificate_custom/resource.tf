# Create a custom certificate
resource "npm_certificate_custom" "example" {
  name = "example.com"

  certificate     = file("example.pem")
  certificate_key = file("example.key")
}
