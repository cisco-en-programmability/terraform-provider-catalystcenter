
data "catalystcenter_wireless_settings_power_profiles_id" "example" {
  provider = catalystcenter
  id       = "string"
}

output "catalystcenter_wireless_settings_power_profiles_id_example" {
  value = data.catalystcenter_wireless_settings_power_profiles_id.example.item
}
