
data "catalystcenter_integration_settings_itsm_instances" "example" {
  provider = catalystcenter
}

output "catalystcenter_integration_settings_itsm_instances_example" {
  value = data.catalystcenter_integration_settings_itsm_instances.example.items
}
