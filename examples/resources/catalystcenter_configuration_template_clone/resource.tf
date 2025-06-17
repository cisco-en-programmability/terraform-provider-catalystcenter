
resource "catalystcenter_configuration_template_clone" "example" {
  provider    = catalystcenter
  name        = "string"
  project_id  = "string"
  template_id = "string"
  parameters = [{

  }]
}

output "catalystcenter_configuration_template_clone_example" {
  value = catalystcenter_configuration_template_clone.example
}