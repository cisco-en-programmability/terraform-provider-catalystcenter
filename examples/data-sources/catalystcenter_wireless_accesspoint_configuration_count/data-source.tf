
data "catalystcenter_wireless_accesspoint_configuration_count" "example" {
  provider       = catalystcenter
  ap_mode        = "string"
  ap_model       = "string"
  mesh_role      = "string"
  provisioned    = "string"
  wlc_ip_address = "string"
}

output "catalystcenter_wireless_accesspoint_configuration_count_example" {
  value = data.catalystcenter_wireless_accesspoint_configuration_count.example.item
}
