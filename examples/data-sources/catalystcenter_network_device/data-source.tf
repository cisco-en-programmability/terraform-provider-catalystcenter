
data "catalystcenter_network_device" "example" {
    provider = catalystcenter
    id = "string"
}

output "catalystcenter_network_device_example" {
    value = data.catalystcenter_network_device.example.item
}
