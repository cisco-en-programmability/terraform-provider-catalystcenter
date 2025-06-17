
data "catalystcenter_network_bugs_results_bugs_id_network_devices_count" "example" {
  provider          = catalystcenter
  id                = "string"
  network_device_id = "string"
  scan_mode         = "string"
  scan_status       = "string"
}

output "catalystcenter_network_bugs_results_bugs_id_network_devices_count_example" {
  value = data.catalystcenter_network_bugs_results_bugs_id_network_devices_count.example.item
}
