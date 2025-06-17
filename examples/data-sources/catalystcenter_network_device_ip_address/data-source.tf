
data "catalystcenter_network_device_ip_address" "example" {
  provider   = catalystcenter
  ip_address = "string"
}

output "catalystcenter_network_device_ip_address_example" {
  value = data.catalystcenter_network_device_ip_address.example.item
}
