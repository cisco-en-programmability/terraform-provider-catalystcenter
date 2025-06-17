
resource "catalystcenter_network_profiles_for_sites_site_assignments" "example" {
  provider = catalystcenter
 
  parameters {

    id         = "string"
    profile_id = "string"
  }
}

output "catalystcenter_network_profiles_for_sites_site_assignments_example" {
  value = catalystcenter_network_profiles_for_sites_site_assignments.example
}