
data "catalystcenter_network_bugs_results_network_devices_network_device_id_bugs_id" "example" {
    provider = catalystcenter
    id = "string"
    network_device_id = "string"
}

output "catalystcenter_network_bugs_results_network_devices_network_device_id_bugs_id_example" {
    value = data.catalystcenter_network_bugs_results_network_devices_network_device_id_bugs_id.example.item
}
