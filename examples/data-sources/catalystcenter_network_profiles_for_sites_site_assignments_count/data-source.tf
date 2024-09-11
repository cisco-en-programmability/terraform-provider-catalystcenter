
data "catalystcenter_network_profiles_for_sites_site_assignments_count" "example" {
  provider   = catalystcenter
  profile_id = "string"
}

output "catalystcenter_network_profiles_for_sites_site_assignments_count_example" {
  value = data.catalystcenter_network_profiles_for_sites_site_assignments_count.example.item
}
