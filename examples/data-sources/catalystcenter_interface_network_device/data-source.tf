
data "catalystcenter_interface_network_device" "example" {
    provider = catalystcenter
    device_id = "string"
}

output "catalystcenter_interface_network_device_example" {
    value = data.catalystcenter_interface_network_device.example.items
}
