
data "catalystcenter_network_device_vlan" "example" {
    provider = catalystcenter
    id = "string"
    interface_type = "string"
}

output "catalystcenter_network_device_vlan_example" {
    value = data.catalystcenter_network_device_vlan.example.items
}
