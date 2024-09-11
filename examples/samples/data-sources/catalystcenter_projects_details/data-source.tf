
data "catalystcenter_projects_details" "example" {
  provider   = catalystcenter
  id         = "string"
  limit      = 1
  name       = "string"
  offset     = 1
  sort_order = "string"
}

output "catalystcenter_projects_details_example" {
  value = data.catalystcenter_projects_details.example.items
}
