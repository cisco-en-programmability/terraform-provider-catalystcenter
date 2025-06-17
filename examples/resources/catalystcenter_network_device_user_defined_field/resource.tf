
resource "catalystcenter_network_device_user_defined_field" "example" {
  provider = catalystcenter
 
  parameters {

    description = "string"
    id          = "string"
    name        = "string"
  }
}

output "catalystcenter_network_device_user_defined_field_example" {
  value = catalystcenter_network_device_user_defined_field.example
}