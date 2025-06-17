
data "catalystcenter_network_bugs_results_network_devices_network_device_id_bugs_count" "example" {
  provider          = catalystcenter
  id                = "string"
  network_device_id = "string"
  severity          = "string"
}

output "catalystcenter_network_bugs_results_network_devices_network_device_id_bugs_count_example" {
  value = data.catalystcenter_network_bugs_results_network_devices_network_device_id_bugs_count.example.item
}
