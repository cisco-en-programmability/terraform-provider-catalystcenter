
data "catalystcenter_sda_fabric_devices" "example" {
    provider = catalystcenter
    device_roles = "string"
    fabric_id = "string"
    limit = 1
    network_device_id = "string"
    offset = 1
}

output "catalystcenter_sda_fabric_devices_example" {
    value = data.catalystcenter_sda_fabric_devices.example.items
}
