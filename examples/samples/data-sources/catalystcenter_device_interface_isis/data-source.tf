
data "catalystcenter_device_interface_isis" "example" {
  provider = catalystcenter
}

output "catalystcenter_device_interface_isis_example" {
  value = data.catalystcenter_device_interface_isis.example.items
}
