
data "catalystcenter_network_device_stack_details" "example" {
    provider = catalystcenter
    device_id = "string"
}

output "catalystcenter_network_device_stack_details_example" {
    value = data.catalystcenter_network_device_stack_details.example.item
}
