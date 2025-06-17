
data "catalystcenter_icap_settings_count" "example" {
  provider       = catalystcenter
  apid           = "string"
  capture_status = "string"
  capture_type   = "string"
  client_mac     = "string"
  wlc_id         = "string"
}

output "catalystcenter_icap_settings_count_example" {
  value = data.catalystcenter_icap_settings_count.example.item
}
