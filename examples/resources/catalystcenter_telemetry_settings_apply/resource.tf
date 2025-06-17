
resource "catalystcenter_telemetry_settings_apply" "example" {
  provider = catalystcenter
  parameters = [{

    device_ids = ["string"]
  }]
}

output "catalystcenter_telemetry_settings_apply_example" {
  value = catalystcenter_telemetry_settings_apply.example
}