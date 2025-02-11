
data "catalystcenter_network_device_maintenance_schedules" "example" {
    provider = catalystcenter
    limit = "string"
    network_device_ids = "string"
    offset = "string"
    order = "string"
    sort_by = "string"
    status = "string"
}

output "catalystcenter_network_device_maintenance_schedules_example" {
    value = data.catalystcenter_network_device_maintenance_schedules.example.items
}
