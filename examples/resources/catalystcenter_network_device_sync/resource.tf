
resource "catalystcenter_network_device_sync" "example" {
  provider   = catalystcenter
  force_sync = "false"
  parameters = ["string"]
}

output "catalystcenter_network_device_sync_example" {
  value = catalystcenter_network_device_sync.example
}