
data "catalystcenter_network_bugs_results_network_devices_network_device_id_bugs" "example" {
  provider          = catalystcenter
  id                = "string"
  limit             = 1
  network_device_id = "string"
  offset            = 1
  order             = "string"
  severity          = "string"
  sort_by           = "string"
}

output "catalystcenter_network_bugs_results_network_devices_network_device_id_bugs_example" {
  value = data.catalystcenter_network_bugs_results_network_devices_network_device_id_bugs.example.items
}
