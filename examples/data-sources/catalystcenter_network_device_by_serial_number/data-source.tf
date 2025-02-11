
data "catalystcenter_network_device_by_serial_number" "example" {
    provider = catalystcenter
    serial_number = "string"
}

output "catalystcenter_network_device_by_serial_number_example" {
    value = data.catalystcenter_network_device_by_serial_number.example.item
}
