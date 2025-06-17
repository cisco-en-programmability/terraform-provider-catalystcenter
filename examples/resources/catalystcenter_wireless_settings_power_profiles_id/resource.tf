
resource "catalystcenter_wireless_settings_power_profiles_id" "example" {
  provider = catalystcenter
 
  parameters {

    description  = "string"
    id           = "string"
    profile_name = "string"
    rules {

      interface_id    = "string"
      interface_type  = "string"
      parameter_type  = "string"
      parameter_value = "string"
    }
  }
}

output "catalystcenter_wireless_settings_power_profiles_id_example" {
  value = catalystcenter_wireless_settings_power_profiles_id.example
}