
data "catalystcenter_wireless_accesspoint_configuration_summary" "example" {
  provider       = catalystcenter
  ap_mode        = "string"
  ap_model       = "string"
  key            = "string"
  limit          = 1
  mesh_role      = "string"
  offset         = 1
  provisioned    = "string"
  wlc_ip_address = "string"
}

output "catalystcenter_wireless_accesspoint_configuration_summary_example" {
  value = data.catalystcenter_wireless_accesspoint_configuration_summary.example.item
}
