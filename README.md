# Laravel Forge Terraform Provider


## Examples

```
provider "forge" {
  token = "your_api_key"
}

data "forge_credential" "credential" {
  name = "test"
}

data "forge_server" "server01" {
  name = "server01"
}

resource "forge_server" "server02" {
    credential_id      = data.forge_credential.credential.id
    name               = "server02"
    php_version        = "php72"
    ip_address         = "66.120.15.15"
    private_ip_address = "10.10.10.10"
    platform           = "ocean2"
    region             = "New York 1"
    size               = "07"
    tags               = ["staging", "dev"]
    network            = [data.forge_server.server01.id]
}

output "ip-address" {
  value = forge_server.server02.ip_address
}

```