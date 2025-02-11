
data "catalystcenter_projects_project_id" "example" {
    provider = catalystcenter
    project_id = "string"
}

output "catalystcenter_projects_project_id_example" {
    value = data.catalystcenter_projects_project_id.example.item
}
