
data "catalystcenter_sites_profile_assignments" "example" {
  provider = catalystcenter
  limit    = 1
  offset   = 1
  site_id  = "string"
}

output "catalystcenter_sites_profile_assignments_example" {
  value = data.catalystcenter_sites_profile_assignments.example.items
}
