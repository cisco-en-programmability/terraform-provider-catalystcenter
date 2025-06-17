
resource "catalystcenter_wireless_profiles" "example" {
  provider = catalystcenter
 
  parameters {

    additional_interfaces = ["string"]
    ap_zones {

      ap_zone_name    = "string"
      rf_profile_name = "string"
      ssids           = ["string"]
    }
    feature_templates {

      id    = "string"
      ssids = ["string"]
    }
    id = "string"
    ssid_details {

      anchor_group_name  = "string"
      dot11be_profile_id = "string"
      enable_fabric      = "false"
      flex_connect {

        enable_flex_connect = "false"
        local_to_vlan       = 1
      }
      interface_name      = "string"
      policy_profile_name = "string"
      ssid_name           = "string"
      vlan_group_name     = "string"
      wlan_profile_name   = "string"
    }
    wireless_profile_name = "string"
  }
}

output "catalystcenter_wireless_profiles_example" {
  value = catalystcenter_wireless_profiles.example
}