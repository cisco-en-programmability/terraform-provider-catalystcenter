
resource "catalystcenter_network_device_update_role" "example" {
  provider = meraki
  parameters {

    id          = "string"
    role        = "string"
    role_source = "string"
  }
}

output "catalystcenter_network_device_update_role_example" {
  value = catalystcenter_network_device_update_role.example
}