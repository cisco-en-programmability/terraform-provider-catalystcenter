
data "catalystcenter_network_device_by_ip" "example" {
  provider   = catalystcenter
  ip_address = "string"
}

output "catalystcenter_network_device_by_ip_example" {
  value = data.catalystcenter_network_device_by_ip.example.item
}
