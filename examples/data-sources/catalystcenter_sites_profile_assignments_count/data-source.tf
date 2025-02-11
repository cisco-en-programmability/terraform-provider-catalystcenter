
data "catalystcenter_sites_profile_assignments_count" "example" {
    provider = catalystcenter
    site_id = "string"
}

output "catalystcenter_sites_profile_assignments_count_example" {
    value = data.catalystcenter_sites_profile_assignments_count.example.item
}
