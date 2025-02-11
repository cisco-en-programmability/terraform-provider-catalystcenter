
data "catalystcenter_wireless_settings_ap_authorization_lists" "example" {
    provider = catalystcenter
    ap_authorization_list_name = "string"
    limit = "string"
    offset = "string"
}

output "catalystcenter_wireless_settings_ap_authorization_lists_example" {
    value = data.catalystcenter_wireless_settings_ap_authorization_lists.example.item
}
