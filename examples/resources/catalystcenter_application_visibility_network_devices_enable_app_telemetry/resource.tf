
resource "catalystcenter_application_visibility_network_devices_enable_app_telemetry" "example" {
  provider = catalystcenter
  parameters = [{

    network_devices = [{

      id                  = "string"
      include_guest_ssids = "false"
      include_wlan_modes  = ["string"]
    }]
  }]
}

output "catalystcenter_application_visibility_network_devices_enable_app_telemetry_example" {
  value = catalystcenter_application_visibility_network_devices_enable_app_telemetry.example
}