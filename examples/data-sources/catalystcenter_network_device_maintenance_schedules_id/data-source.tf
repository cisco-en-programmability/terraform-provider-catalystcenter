
data "catalystcenter_network_device_maintenance_schedules_id" "example" {
  provider = catalystcenter
  id       = "string"
}

output "catalystcenter_network_device_maintenance_schedules_id_example" {
  value = data.catalystcenter_network_device_maintenance_schedules_id.example.item
}
