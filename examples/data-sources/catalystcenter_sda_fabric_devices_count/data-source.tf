
data "catalystcenter_sda_fabric_devices_count" "example" {
  provider          = catalystcenter
  device_roles      = "string"
  fabric_id         = "string"
  network_device_id = "string"
}

output "catalystcenter_sda_fabric_devices_count_example" {
  value = data.catalystcenter_sda_fabric_devices_count.example.item
}
