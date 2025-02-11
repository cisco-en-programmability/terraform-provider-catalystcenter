
data "catalystcenter_interface_network_device_range" "example" {
    provider = catalystcenter
    device_id = "string"
    records_to_return = 1
    start_index = 1
}

output "catalystcenter_interface_network_device_range_example" {
    value = data.catalystcenter_interface_network_device_range.example.items
}
