
data "catalystcenter_field_notices_results_notices_id_network_devices_count" "example" {
    provider = catalystcenter
    id = "string"
    network_device_id = "string"
    scan_status = "string"
}

output "catalystcenter_field_notices_results_notices_id_network_devices_count_example" {
    value = data.catalystcenter_field_notices_results_notices_id_network_devices_count.example.item
}
