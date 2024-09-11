
resource "catalystcenter_auth_token_create" "example" {
  provider      = catalystcenter
  authorization = "string"
}

output "catalystcenter_auth_token_create_example" {
  value = catalystcenter_auth_token_create.example
}