
resource "catalystcenter_network_profiles_for_sites_site_assignments_bulk_delete" "example" {
  provider   = meraki
  profile_id = "string"
  site_id    = "string"
  parameters {

  }
}

output "catalystcenter_network_profiles_for_sites_site_assignments_bulk_delete_example" {
  value = catalystcenter_network_profiles_for_sites_site_assignments_bulk_delete.example
}