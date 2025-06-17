
resource "catalystcenter_network_device_user_defined_field_delete" "example" {
  provider  = catalystcenter
  device_id = "string"
  name      = "string"
  parameters = [{

  }]
}

output "catalystcenter_network_device_user_defined_field_delete_example" {
  value = catalystcenter_network_device_user_defined_field_delete.example
}