
resource "catalystcenter_wireless_profiles" "example" {
  provider = catalystcenter

  parameters {

    id = "string"
    ssid_details {

      dot11be_profile_id = "string"
      enable_fabric      = "false"
      flex_connect {

        enable_flex_connect = "false"
        local_to_vlan       = 1
      }
      interface_name    = "string"
      ssid_name         = "string"
      wlan_profile_name = "string"
    }
    wireless_profile_name = "string"
  }
}

output "catalystcenter_wireless_profiles_example" {
  value = catalystcenter_wireless_profiles.example
}