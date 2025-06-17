
data "catalystcenter_field_notices_results_network_devices_network_device_id_notices_count" "example" {
  provider          = catalystcenter
  id                = "string"
  network_device_id = "string"
  type              = "string"
}

output "catalystcenter_field_notices_results_network_devices_network_device_id_notices_count_example" {
  value = data.catalystcenter_field_notices_results_network_devices_network_device_id_notices_count.example.item
}
