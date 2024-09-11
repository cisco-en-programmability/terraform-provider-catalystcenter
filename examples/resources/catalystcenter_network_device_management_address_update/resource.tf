
resource "catalystcenter_network_device_management_address_update" "example" {
  provider = catalystcenter
  deviceid = "string"
  parameters {

    new_ip = "string"
  }
}

output "catalystcenter_network_device_management_address_update_example" {
  value = catalystcenter_network_device_management_address_update.example
}