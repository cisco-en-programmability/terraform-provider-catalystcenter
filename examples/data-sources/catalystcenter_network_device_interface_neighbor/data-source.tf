
data "catalystcenter_network_device_interface_neighbor" "example" {
    provider = catalystcenter
    device_uuid = "string"
    interface_uuid = "string"
}

output "catalystcenter_network_device_interface_neighbor_example" {
    value = data.catalystcenter_network_device_interface_neighbor.example.item
}
