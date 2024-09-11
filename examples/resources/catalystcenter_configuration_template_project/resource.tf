
resource "catalystcenter_configuration_template_project" "example" {
  provider = catalystcenter

  parameters {

    create_time      = 1
    description      = "string"
    id               = "string"
    last_update_time = 1
    name             = "string"
    project_id       = "string"
    tags {

      id   = "string"
      name = "string"
    }
  }
}

output "catalystcenter_configuration_template_project_example" {
  value = catalystcenter_configuration_template_project.example
}