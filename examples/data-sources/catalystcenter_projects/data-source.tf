
data "catalystcenter_projects" "example" {
  provider = catalystcenter
  limit    = 1
  name     = "string"
  offset   = 1
}

output "catalystcenter_projects_example" {
  value = data.catalystcenter_projects.example.items
}
