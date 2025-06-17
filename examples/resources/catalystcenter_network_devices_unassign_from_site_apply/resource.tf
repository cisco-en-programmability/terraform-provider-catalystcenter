
resource "catalystcenter_network_devices_unassign_from_site_apply" "example" {
  provider = catalystcenter
  parameters = [{

    device_ids = ["string"]
  }]
}

output "catalystcenter_network_devices_unassign_from_site_apply_example" {
  value = catalystcenter_network_devices_unassign_from_site_apply.example
}