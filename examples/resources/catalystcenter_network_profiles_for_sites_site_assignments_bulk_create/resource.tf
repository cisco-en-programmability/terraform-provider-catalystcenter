
resource "catalystcenter_network_profiles_for_sites_site_assignments_bulk_create" "example" {
  provider   = meraki
  profile_id = "string"
  parameters {

    type = "------"
  }
}

output "catalystcenter_network_profiles_for_sites_site_assignments_bulk_create_example" {
  value = catalystcenter_network_profiles_for_sites_site_assignments_bulk_create.example
}