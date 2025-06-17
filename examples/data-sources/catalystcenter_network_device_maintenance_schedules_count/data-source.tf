
data "catalystcenter_network_device_maintenance_schedules_count" "example" {
  provider           = catalystcenter
  network_device_ids = "string"
  status             = "string"
}

output "catalystcenter_network_device_maintenance_schedules_count_example" {
  value = data.catalystcenter_network_device_maintenance_schedules_count.example.item
}
