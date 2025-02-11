
data "catalystcenter_wireless_settings_power_profiles" "example" {
    provider = catalystcenter
    limit = 1
    offset = 1
    profile_name = "string"
}

output "catalystcenter_wireless_settings_power_profiles_example" {
    value = data.catalystcenter_wireless_settings_power_profiles.example.items
}
