
data "catalystcenter_security_advisories_results_advisories_id_network_devices_network_device_id" "example" {
    provider = catalystcenter
    id = "string"
    network_device_id = "string"
}

output "catalystcenter_security_advisories_results_advisories_id_network_devices_network_device_id_example" {
    value = data.catalystcenter_security_advisories_results_advisories_id_network_devices_network_device_id.example.item
}
