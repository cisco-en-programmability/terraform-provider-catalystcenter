
data "catalystcenter_security_advisories_results_network_devices_network_device_id_advisories_count" "example" {
  provider               = catalystcenter
  cvss_base_score        = "string"
  id                     = "string"
  network_device_id      = "string"
  security_impact_rating = "string"
}

output "catalystcenter_security_advisories_results_network_devices_network_device_id_advisories_count_example" {
  value = data.catalystcenter_security_advisories_results_network_devices_network_device_id_advisories_count.example.item
}
