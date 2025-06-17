
data "catalystcenter_network_devices_resync_interval_settings_id" "example" {
  provider = catalystcenter
  id       = "string"
}

output "catalystcenter_network_devices_resync_interval_settings_id_example" {
  value = data.catalystcenter_network_devices_resync_interval_settings_id.example.item
}
