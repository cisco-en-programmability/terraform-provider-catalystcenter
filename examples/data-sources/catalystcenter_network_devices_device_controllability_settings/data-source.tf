
data "catalystcenter_network_devices_device_controllability_settings" "example" {
    provider = catalystcenter
}

output "catalystcenter_network_devices_device_controllability_settings_example" {
    value = data.catalystcenter_network_devices_device_controllability_settings.example.item
}
