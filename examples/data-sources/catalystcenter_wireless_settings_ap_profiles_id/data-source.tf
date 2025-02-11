
data "catalystcenter_wireless_settings_ap_profiles_id" "example" {
    provider = catalystcenter
    id = "string"
}

output "catalystcenter_wireless_settings_ap_profiles_id_example" {
    value = data.catalystcenter_wireless_settings_ap_profiles_id.example.items
}
