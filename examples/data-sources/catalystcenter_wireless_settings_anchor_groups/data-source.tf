
data "catalystcenter_wireless_settings_anchor_groups" "example" {
    provider = catalystcenter
}

output "catalystcenter_wireless_settings_anchor_groups_example" {
    value = data.catalystcenter_wireless_settings_anchor_groups.example.item
}
