
data "catalystcenter_sites_telemetry_settings" "example" {
  provider  = catalystcenter
  id        = "string"
  inherited = "false"
}

output "catalystcenter_sites_telemetry_settings_example" {
  value = data.catalystcenter_sites_telemetry_settings.example.item
}
