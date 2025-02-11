
resource "catalystcenter_network_devices_resync_interval_settings_override" "example" {
    provider = meraki
}

output "catalystcenter_network_devices_resync_interval_settings_override_example" {
    value = catalystcenter_network_devices_resync_interval_settings_override.example
}