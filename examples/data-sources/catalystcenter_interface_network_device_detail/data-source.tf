
data "catalystcenter_interface_network_device_detail" "example" {
    provider = catalystcenter
    device_id = "string"
    name = "string"
}

output "catalystcenter_interface_network_device_detail_example" {
    value = data.catalystcenter_interface_network_device_detail.example.item
}
