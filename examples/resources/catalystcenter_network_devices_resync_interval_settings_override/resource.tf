
resource "catalystcenter_network_devices_resync_interval_settings_override" "example" {
  provider = catalystcenter
}

output "catalystcenter_network_devices_resync_interval_settings_override_example" {
  value = catalystcenter_network_devices_resync_interval_settings_override.example
}