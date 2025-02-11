
data "catalystcenter_network_device_chassis_details" "example" {
    provider = catalystcenter
    device_id = "string"
}

output "catalystcenter_network_device_chassis_details_example" {
    value = data.catalystcenter_network_device_chassis_details.example.items
}
