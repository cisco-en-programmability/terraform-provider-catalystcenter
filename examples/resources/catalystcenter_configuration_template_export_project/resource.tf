
resource "catalystcenter_configuration_template_export_project" "example" {
    provider = meraki
    parameters = ["string"]
}

output "catalystcenter_configuration_template_export_project_example" {
    value = catalystcenter_configuration_template_export_project.example
}