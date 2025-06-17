
data "catalystcenter_network_device_interface_poe" "example" {
  provider            = catalystcenter
  device_uuid         = "string"
  interface_name_list = "string"
}

output "catalystcenter_network_device_interface_poe_example" {
  value = data.catalystcenter_network_device_interface_poe.example.items
}
