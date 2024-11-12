
resource "catalystcenter_users_external_servers_aaa_attribute" "example" {
  provider = catalystcenter


  parameters {

    attribute_name = "string"
  }
}

output "catalystcenter_users_external_servers_aaa_attribute_example" {
  value = catalystcenter_users_external_servers_aaa_attribute.example
}