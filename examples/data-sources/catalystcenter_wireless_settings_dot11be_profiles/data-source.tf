
data "catalystcenter_wireless_settings_dot11be_profiles" "example" {
  provider = catalystcenter
  limit    = 1
  offset   = 1
}

output "catalystcenter_wireless_settings_dot11be_profiles_example" {
  value = data.catalystcenter_wireless_settings_dot11be_profiles.example.items
}

data "catalystcenter_wireless_settings_dot11be_profiles" "example" {
  provider = catalystcenter
  id       = "string"
}

output "catalystcenter_wireless_settings_dot11be_profiles_example" {
  value = data.catalystcenter_wireless_settings_dot11be_profiles.example.item
}
