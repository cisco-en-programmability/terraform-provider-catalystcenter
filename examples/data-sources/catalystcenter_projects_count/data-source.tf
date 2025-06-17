
data "catalystcenter_projects_count" "example" {
  provider = catalystcenter
  name     = "string"
}

output "catalystcenter_projects_count_example" {
  value = data.catalystcenter_projects_count.example.item
}
