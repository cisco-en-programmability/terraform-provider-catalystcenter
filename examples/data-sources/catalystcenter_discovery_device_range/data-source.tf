
data "catalystcenter_discovery_device_range" "example" {
    provider = catalystcenter
    id = "string"
    records_to_return = 1
    start_index = 1
    task_id = "string"
}

output "catalystcenter_discovery_device_range_example" {
    value = data.catalystcenter_discovery_device_range.example.items
}
