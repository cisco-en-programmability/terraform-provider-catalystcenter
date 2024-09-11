
data "catalystcenter_integration_settings_status" "example" {
  provider = catalystcenter
}

output "catalystcenter_integration_settings_status_example" {
  value = data.catalystcenter_integration_settings_status.example.items
}
