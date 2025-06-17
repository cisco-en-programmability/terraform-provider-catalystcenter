
data "catalystcenter_wireless_settings_rf_profiles" "example" {
  provider                = catalystcenter
  enable_radio_type6_g_hz = "false"
  enable_radio_type_a     = "false"
  enable_radio_type_b     = "false"
  limit                   = 1
  offset                  = 1
  rf_profile_name         = "string"
}

output "catalystcenter_wireless_settings_rf_profiles_example" {
  value = data.catalystcenter_wireless_settings_rf_profiles.example.items
}

data "catalystcenter_wireless_settings_rf_profiles" "example" {
  provider = catalystcenter
  id       = "string"
}

output "catalystcenter_wireless_settings_rf_profiles_example" {
  value = data.catalystcenter_wireless_settings_rf_profiles.example.item
}
