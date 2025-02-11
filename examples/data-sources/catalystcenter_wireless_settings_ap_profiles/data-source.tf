
data "catalystcenter_wireless_settings_ap_profiles" "example" {
    provider = catalystcenter
    ap_profile_name = "string"
    limit = "string"
    offset = "string"
}

output "catalystcenter_wireless_settings_ap_profiles_example" {
    value = data.catalystcenter_wireless_settings_ap_profiles.example.items
}
