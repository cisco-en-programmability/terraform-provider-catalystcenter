
data "catalystcenter_network_device_user_defined_field" "example" {
  provider = catalystcenter
  id       = "string"
  name     = "string"
}

output "catalystcenter_network_device_user_defined_field_example" {
  value = data.catalystcenter_network_device_user_defined_field.example.items
}
