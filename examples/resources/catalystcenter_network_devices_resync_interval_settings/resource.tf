
resource "catalystcenter_network_devices_resync_interval_settings" "example" {
  provider = meraki
  parameters {

    interval = 1
  }
}

output "catalystcenter_network_devices_resync_interval_settings_example" {
  value = catalystcenter_network_devices_resync_interval_settings.example
}