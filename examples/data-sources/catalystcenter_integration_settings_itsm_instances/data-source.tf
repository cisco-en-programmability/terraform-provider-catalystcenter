
data "catalystcenter_integration_settings_itsm_instances" "example" {
  provider  = catalystcenter
  order     = "string"
  page      = 1.0
  page_size = 1.0
  sort_by   = "string"
}

output "catalystcenter_integration_settings_itsm_instances_example" {
  value = data.catalystcenter_integration_settings_itsm_instances.example.item
}
