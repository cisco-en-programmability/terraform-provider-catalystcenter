
resource "catalystcenter_network_devices_resync_interval_settings_id" "example" {
  provider = catalystcenter

  parameters {

    id       = "string"
    interval = 1
  }
}

output "catalystcenter_network_devices_resync_interval_settings_id_example" {
  value = catalystcenter_network_devices_resync_interval_settings_id.example
}