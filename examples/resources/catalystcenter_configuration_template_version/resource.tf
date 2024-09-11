
resource "catalystcenter_configuration_template_version" "example" {
  provider = catalystcenter
  parameters {
    # comments = "string"
    template_id = "string"
  }
}

output "catalystcenter_configuration_template_version_example" {
  value = catalystcenter_configuration_template_version.example
}