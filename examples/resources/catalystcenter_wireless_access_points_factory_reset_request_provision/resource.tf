
resource "catalystcenter_wireless_access_points_factory_reset_request_provision" "example" {
  provider = catalystcenter
  parameters {

    ap_mac_addresses     = ["string"]
    keep_static_ipconfig = "false"
  }
}

output "catalystcenter_wireless_access_points_factory_reset_request_provision_example" {
  value = catalystcenter_wireless_access_points_factory_reset_request_provision.example
}