
resource "catalystcenter_network_devices_device_controllability_settings" "example" {
  provider = catalystcenter
 
  parameters {

    autocorrect_telemetry_config = "false"
    device_controllability       = "false"
  }
}

output "catalystcenter_network_devices_device_controllability_settings_example" {
  value = catalystcenter_network_devices_device_controllability_settings.example
}