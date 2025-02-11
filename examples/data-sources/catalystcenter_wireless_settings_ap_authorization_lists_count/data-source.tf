
data "catalystcenter_wireless_settings_ap_authorization_lists_count" "example" {
    provider = catalystcenter
}

output "catalystcenter_wireless_settings_ap_authorization_lists_count_example" {
    value = data.catalystcenter_wireless_settings_ap_authorization_lists_count.example.item
}
