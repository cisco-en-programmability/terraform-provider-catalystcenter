
resource "catalystcenter_wireless_controllers_assign_managed_ap_locations" "example" {
  provider  = catalystcenter
  device_id = "string"
  parameters {

    primary_managed_aplocations_site_ids   = ["string"]
    secondary_managed_aplocations_site_ids = ["string"]
  }
}

output "catalystcenter_wireless_controllers_assign_managed_ap_locations_example" {
  value = catalystcenter_wireless_controllers_assign_managed_ap_locations.example
}