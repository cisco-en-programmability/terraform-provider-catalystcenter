
data "catalystcenter_field_notices_results_network_devices_count" "example" {
    provider = catalystcenter
    network_device_id = "string"
    notice_count = 1.0
    scan_status = "string"
}

output "catalystcenter_field_notices_results_network_devices_count_example" {
    value = data.catalystcenter_field_notices_results_network_devices_count.example.item
}
