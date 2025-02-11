
data "catalystcenter_network_profiles_for_sites_site_assignments" "example" {
    provider = catalystcenter
    limit = 1
    offset = 1
    profile_id = "string"
}

output "catalystcenter_network_profiles_for_sites_site_assignments_example" {
    value = data.catalystcenter_network_profiles_for_sites_site_assignments.example.items
}
