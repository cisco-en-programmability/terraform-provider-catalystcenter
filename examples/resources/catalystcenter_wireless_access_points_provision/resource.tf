
resource "catalystcenter_wireless_access_points_provision" "example" {
  provider = meraki
  parameters {

    ap_zone_name = "string"
    network_devices {

      device_id = "string"
      mesh_role = "string"
    }
    rf_profile_name = "string"
    site_id         = "string"
  }
}

output "catalystcenter_wireless_access_points_provision_example" {
  value = catalystcenter_wireless_access_points_provision.example
}