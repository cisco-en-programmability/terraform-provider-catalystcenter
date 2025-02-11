
resource "catalystcenter_projects_project_id" "example" {
    provider = catalystcenter

    parameters {

      description = "string"
      name = "string"
      project_id = "string"
    }
}

output "catalystcenter_projects_project_id_example" {
    value = catalystcenter_projects_project_id.example
}
