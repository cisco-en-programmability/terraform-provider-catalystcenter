
data "catalystcenter_security_advisories_results_network_devices_network_device_id" "example" {
    provider = catalystcenter
    network_device_id = "string"
}

output "catalystcenter_security_advisories_results_network_devices_network_device_id_example" {
    value = data.catalystcenter_security_advisories_results_network_devices_network_device_id.example.item
}
