
data "catalystcenter_wireless_settings_anchor_groups_count" "example" {
    provider = catalystcenter
}

output "catalystcenter_wireless_settings_anchor_groups_count_example" {
    value = data.catalystcenter_wireless_settings_anchor_groups_count.example.item
}
