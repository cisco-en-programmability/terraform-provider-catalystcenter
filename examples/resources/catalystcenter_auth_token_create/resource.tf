
resource "catalystcenter_auth_token_create" "example" {
  provider      = meraki
  authorization = "string"
}

output "catalystcenter_auth_token_create_example" {
  value = catalystcenter_auth_token_create.example
}