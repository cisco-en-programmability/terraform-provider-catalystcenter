
data "catalystcenter_sda_fabric_edge_device" "example" {
    provider = catalystcenter
    device_management_ip_address = "string"
}

output "catalystcenter_sda_fabric_edge_device_example" {
    value = data.catalystcenter_sda_fabric_edge_device.example.item
}
