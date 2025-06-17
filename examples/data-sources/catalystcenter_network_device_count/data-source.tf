
data "catalystcenter_network_device_count" "example" {
  provider  = catalystcenter
  device_id = "string"
}

output "catalystcenter_network_device_count_example" {
  value = data.catalystcenter_network_device_count.example.item
}
