
data "catalystcenter_network_profiles_for_sites_profile_id_templates_count" "example" {
  provider   = catalystcenter
  profile_id = "string"
}

output "catalystcenter_network_profiles_for_sites_profile_id_templates_count_example" {
  value = data.catalystcenter_network_profiles_for_sites_profile_id_templates_count.example.item
}
