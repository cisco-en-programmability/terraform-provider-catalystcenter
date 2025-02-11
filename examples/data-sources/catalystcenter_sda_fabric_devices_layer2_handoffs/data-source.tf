
data "catalystcenter_sda_fabric_devices_layer2_handoffs" "example" {
    provider = catalystcenter
    fabric_id = "string"
    limit = 1
    network_device_id = "string"
    offset = 1
}

output "catalystcenter_sda_fabric_devices_layer2_handoffs_example" {
    value = data.catalystcenter_sda_fabric_devices_layer2_handoffs.example.items
}
