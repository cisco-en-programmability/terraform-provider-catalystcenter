
resource "catalystcenter_configuration_template_export_template" "example" {
  provider   = catalystcenter
  parameters = ["string"]
}

output "catalystcenter_configuration_template_export_template_example" {
  value = catalystcenter_configuration_template_export_template.example
}