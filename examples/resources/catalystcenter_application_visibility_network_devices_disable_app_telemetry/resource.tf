
resource "catalystcenter_application_visibility_network_devices_disable_app_telemetry" "example" {
  provider = catalystcenter
  parameters = [{

    network_device_ids = ["string"]
  }]
}

output "catalystcenter_application_visibility_network_devices_disable_app_telemetry_example" {
  value = catalystcenter_application_visibility_network_devices_disable_app_telemetry.example
}