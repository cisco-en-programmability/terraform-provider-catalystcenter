
data "catalystcenter_wireless_settings_anchor_groups_id" "example" {
  provider = catalystcenter
  id       = "string"
}

output "catalystcenter_wireless_settings_anchor_groups_id_example" {
  value = data.catalystcenter_wireless_settings_anchor_groups_id.example.item
}
