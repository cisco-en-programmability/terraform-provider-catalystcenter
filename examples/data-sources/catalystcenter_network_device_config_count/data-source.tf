
data "catalystcenter_network_device_config_count" "example" {
  provider = catalystcenter
}

output "catalystcenter_network_device_config_count_example" {
  value = data.catalystcenter_network_device_config_count.example.item
}
