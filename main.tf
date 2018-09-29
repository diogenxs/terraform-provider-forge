provider "forge" {
  token = "insert_api_key_here"
}

resource "forge_server" "web" {
  platform    = "digitalocean"
  name        = "test-web-01"
  region      = "nyc3"
  size        = "4gb"
  php_version = "php72"
}
