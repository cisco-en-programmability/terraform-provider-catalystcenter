
resource "catalystcenter_users_external_authentication_create" "example" {
  provider = catalystcenter
  parameters = [{

    enable = "false"
  }]
}

output "catalystcenter_users_external_authentication_create_example" {
  value = catalystcenter_users_external_authentication_create.example
}