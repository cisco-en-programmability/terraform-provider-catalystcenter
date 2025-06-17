
resource "catalystcenter_wireless_settings_network_device_id_assign_anchor_managed_ap_locations" "example" {
  provider          = catalystcenter
  network_device_id = "string"
  parameters = [{

    anchor_managed_aplocations_site_ids = ["string"]
  }]
}

output "catalystcenter_wireless_settings_network_device_id_assign_anchor_managed_ap_locations_example" {
  value = catalystcenter_wireless_settings_network_device_id_assign_anchor_managed_ap_locations.example
}