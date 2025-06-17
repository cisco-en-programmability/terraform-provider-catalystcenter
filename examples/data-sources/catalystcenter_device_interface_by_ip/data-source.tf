
data "catalystcenter_device_interface_by_ip" "example" {
  provider   = catalystcenter
  ip_address = "string"
}

output "catalystcenter_device_interface_by_ip_example" {
  value = data.catalystcenter_device_interface_by_ip.example.items
}
