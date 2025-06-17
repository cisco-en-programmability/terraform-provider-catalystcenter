
data "catalystcenter_configuration_template_version" "example" {
  provider    = catalystcenter
  template_id = "string"
}

output "catalystcenter_configuration_template_version_example" {
  value = data.catalystcenter_configuration_template_version.example.items
}
