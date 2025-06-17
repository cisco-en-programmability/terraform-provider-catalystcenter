
resource "catalystcenter_application_visibility_network_devices_disable_cbar" "example" {
  provider = catalystcenter
  parameters = [{

    network_device_ids = ["string"]
  }]
}

output "catalystcenter_application_visibility_network_devices_disable_cbar_example" {
  value = catalystcenter_application_visibility_network_devices_disable_cbar.example
}