
data "catalystcenter_network_device_poe" "example" {
    provider = catalystcenter
    device_uuid = "string"
}

output "catalystcenter_network_device_poe_example" {
    value = data.catalystcenter_network_device_poe.example.item
}
