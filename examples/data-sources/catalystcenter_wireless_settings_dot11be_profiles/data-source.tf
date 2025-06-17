
data "catalystcenter_wireless_settings_dot11be_profiles" "example" {
  provider             = catalystcenter
  is_mu_mimo_down_link = "false"
  is_mu_mimo_up_link   = "false"
  is_of_dma_down_link  = "false"
  is_of_dma_multi_ru   = "false"
  is_of_dma_up_link    = "false"
  limit                = 1
  offset               = 1
  profile_name         = "string"
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
