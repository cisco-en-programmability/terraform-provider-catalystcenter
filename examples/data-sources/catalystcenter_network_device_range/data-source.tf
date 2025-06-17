
data "catalystcenter_network_device_range" "example" {
  provider          = catalystcenter
  records_to_return = 1
  start_index       = 1
}

output "catalystcenter_network_device_range_example" {
  value = data.catalystcenter_network_device_range.example.items
}
