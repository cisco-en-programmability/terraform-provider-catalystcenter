
data "catalystcenter_wireless_settings_ap_profiles_count" "example" {
  provider = catalystcenter
}

output "catalystcenter_wireless_settings_ap_profiles_count_example" {
  value = data.catalystcenter_wireless_settings_ap_profiles_count.example.item
}
