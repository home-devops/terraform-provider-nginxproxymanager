# Configuration-based authentication
provider "npm" {
  url      = "http://localhost:81"
  username = "admin@example.com"
  password = "password"
}

# Environment variable-based authentication
provider "npm" {}
