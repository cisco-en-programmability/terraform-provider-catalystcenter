
resource "catalystcenter_wireless_settings_dot11be_profiles" "example" {
  provider = catalystcenter
 
  parameters {

    id                = "string"
    mu_mimo_down_link = "false"
    mu_mimo_up_link   = "false"
    ofdma_down_link   = "false"
    ofdma_multi_ru    = "false"
    ofdma_up_link     = "false"
    profile_name      = "string"
  }
}

output "catalystcenter_wireless_settings_dot11be_profiles_example" {
  value = catalystcenter_wireless_settings_dot11be_profiles.example
}