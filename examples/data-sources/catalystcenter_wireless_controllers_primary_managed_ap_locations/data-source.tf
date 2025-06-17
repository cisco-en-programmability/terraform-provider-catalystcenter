
data "catalystcenter_wireless_controllers_primary_managed_ap_locations" "example" {
  provider          = catalystcenter
  limit             = 1
  network_device_id = "string"
  offset            = 1
}

output "catalystcenter_wireless_controllers_primary_managed_ap_locations_example" {
  value = data.catalystcenter_wireless_controllers_primary_managed_ap_locations.example.item
}
