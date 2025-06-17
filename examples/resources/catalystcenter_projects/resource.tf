
resource "catalystcenter_projects" "example" {
  provider = catalystcenter
 
  parameters {

    description = "string"
    name        = "string"
  }
}

output "catalystcenter_projects_example" {
  value = catalystcenter_projects.example
}