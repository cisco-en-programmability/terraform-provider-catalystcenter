
resource "catalystcenter_sda_authentication_profiles" "example" {
  provider = catalystcenter
 
  parameters {

    authentication_order          = "string"
    authentication_profile_name   = "string"
    dot1x_to_mab_fallback_timeout = 1
    fabric_id                     = "string"
    id                            = "string"
    is_bpdu_guard_enabled         = "false"
    is_voice_vlan_enabled         = "false"
    number_of_hosts               = "string"
    pre_auth_acl {

      access_contracts {

        action   = "string"
        port     = "string"
        protocol = "string"
      }
      description     = "string"
      enabled         = "false"
      implicit_action = "string"
    }
    wake_on_lan = "false"
  }
}

output "catalystcenter_sda_authentication_profiles_example" {
  value = catalystcenter_sda_authentication_profiles.example
}