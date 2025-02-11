
data "catalystcenter_sda_fabric_devices_layer2_handoffs_count" "example" {
    provider = catalystcenter
    fabric_id = "string"
    network_device_id = "string"
}

output "catalystcenter_sda_fabric_devices_layer2_handoffs_count_example" {
    value = data.catalystcenter_sda_fabric_devices_layer2_handoffs_count.example.item
}
