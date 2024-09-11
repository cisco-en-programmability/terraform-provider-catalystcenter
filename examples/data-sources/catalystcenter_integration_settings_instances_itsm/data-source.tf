
data "catalystcenter_integration_settings_instances_itsm" "example" {
  provider    = catalystcenter
  instance_id = "string"
}

output "catalystcenter_integration_settings_instances_itsm_example" {
  value = data.catalystcenter_integration_settings_instances_itsm.example.item
}
