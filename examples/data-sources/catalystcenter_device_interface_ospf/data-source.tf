
data "catalystcenter_device_interface_ospf" "example" {
    provider = catalystcenter
}

output "catalystcenter_device_interface_ospf_example" {
    value = data.catalystcenter_device_interface_ospf.example.items
}
