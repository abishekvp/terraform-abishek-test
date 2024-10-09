terraform {
  required_providers {
    abiraj = {
      source = "abirajvp/abiraj"
      version = "0.2.0"
    }
  }
}

provider "abiraj" {
    account_id = "2000000004406"
    server_url = "http://localhost:8000"
    authtoken = "d3753a91-1d93-4e0f-be2e-7f3b2aa8d0c5"
}

data "abiraj_keyvalue" "name" {}

output "password" {
  value = data.abiraj_keyvalue.name.password
}