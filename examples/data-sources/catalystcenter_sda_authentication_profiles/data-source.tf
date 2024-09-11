
data "catalystcenter_sda_authentication_profiles" "example" {
  provider                    = catalystcenter
  authentication_profile_name = "string"
  fabric_id                   = "string"
  limit                       = 1
  offset                      = 1
}

output "catalystcenter_sda_authentication_profiles_example" {
  value = data.catalystcenter_sda_authentication_profiles.example.items
}
