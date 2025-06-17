
data "catalystcenter_network_device_global_polling_interval" "example" {
  provider = catalystcenter
}

output "catalystcenter_network_device_global_polling_interval_example" {
  value = data.catalystcenter_network_device_global_polling_interval.example.item
}
