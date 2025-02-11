
data "catalystcenter_network_device_equipment" "example" {
    provider = catalystcenter
    device_uuid = "string"
    type = "string"
}

output "catalystcenter_network_device_equipment_example" {
    value = data.catalystcenter_network_device_equipment.example.items
}
