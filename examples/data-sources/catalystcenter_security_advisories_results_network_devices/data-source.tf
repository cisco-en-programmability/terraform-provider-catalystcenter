
data "catalystcenter_security_advisories_results_network_devices" "example" {
    provider = catalystcenter
    advisory_count = "string"
    limit = 1
    network_device_id = "string"
    offset = 1
    order = "string"
    scan_mode = "string"
    scan_status = "string"
    sort_by = "string"
}

output "catalystcenter_security_advisories_results_network_devices_example" {
    value = data.catalystcenter_security_advisories_results_network_devices.example.items
}
