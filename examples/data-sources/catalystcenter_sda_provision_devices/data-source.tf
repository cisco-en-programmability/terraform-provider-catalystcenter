
data "catalystcenter_sda_provision_devices" "example" {
    provider = catalystcenter
    id = "string"
    limit = 1
    network_device_id = "string"
    offset = 1
    site_id = "string"
}

output "catalystcenter_sda_provision_devices_example" {
    value = data.catalystcenter_sda_provision_devices.example.items
}
