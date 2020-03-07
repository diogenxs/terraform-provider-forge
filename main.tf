provider "forge" {
  token = "insert_api_key_here"
}

data "forge_credential" "test" {
  name = "test"
}

resource "forge_server" "web" {
  platform      = "digitalocean"
  name          = "test-web-01"
  region        = "nyc3"
  size          = "4gb"
  php_version   = "php72"
  credential_id = data.forge_credential.test.id
}
