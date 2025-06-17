
data "catalystcenter_wireless_settings_ap_authorization_lists_id" "example" {
  provider = catalystcenter
  id       = "string"
}

output "catalystcenter_wireless_settings_ap_authorization_lists_id_example" {
  value = data.catalystcenter_wireless_settings_ap_authorization_lists_id.example.item
}
