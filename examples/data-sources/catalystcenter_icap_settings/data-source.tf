
data "catalystcenter_icap_settings" "example" {
  provider       = catalystcenter
  apid           = "string"
  capture_status = "string"
  capture_type   = "string"
  client_mac     = "string"
  limit          = 1
  offset         = 1
  wlc_id         = "string"
}

output "catalystcenter_icap_settings_example" {
  value = data.catalystcenter_icap_settings.example.items
}
