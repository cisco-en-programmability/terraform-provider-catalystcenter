
data "catalystcenter_network_profiles_for_sites_profile_id_templates" "example" {
  provider   = catalystcenter
  profile_id = "string"
}

output "catalystcenter_network_profiles_for_sites_profile_id_templates_example" {
  value = data.catalystcenter_network_profiles_for_sites_profile_id_templates.example.items
}
