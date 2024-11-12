
resource "catalystcenter_network_device_user_defined_field_update" "example" {
  provider  = meraki
  device_id = "string"
  parameters {

    name  = "string"
    value = "string"
  }
}

output "catalystcenter_network_device_user_defined_field_update_example" {
  value = catalystcenter_network_device_user_defined_field_update.example
}