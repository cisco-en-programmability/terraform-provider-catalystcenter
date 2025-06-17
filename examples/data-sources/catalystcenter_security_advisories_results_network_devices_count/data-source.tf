
data "catalystcenter_security_advisories_results_network_devices_count" "example" {
  provider          = catalystcenter
  advisory_count    = 1.0
  network_device_id = "string"
  scan_mode         = "string"
  scan_status       = "string"
}

output "catalystcenter_security_advisories_results_network_devices_count_example" {
  value = data.catalystcenter_security_advisories_results_network_devices_count.example.item
}
