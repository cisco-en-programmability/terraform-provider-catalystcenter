
resource "catalystcenter_configuration_template_import_project" "example" {
  provider   = catalystcenter
  do_version = "false"
}

output "catalystcenter_configuration_template_import_project_example" {
  value = catalystcenter_configuration_template_import_project.example
}