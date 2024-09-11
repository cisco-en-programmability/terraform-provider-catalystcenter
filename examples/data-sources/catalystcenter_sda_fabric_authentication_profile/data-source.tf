
data "catalystcenter_sda_fabric_authentication_profile" "example" {
  provider                   = catalystcenter
  authenticate_template_name = "string"
  site_name_hierarchy        = "string"
}

output "catalystcenter_sda_fabric_authentication_profile_example" {
  value = data.catalystcenter_sda_fabric_authentication_profile.example.item
}
