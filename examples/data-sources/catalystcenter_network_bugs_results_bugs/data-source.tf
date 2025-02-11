
data "catalystcenter_network_bugs_results_bugs" "example" {
    provider = catalystcenter
    device_count = 1.0
    id = "string"
    limit = 1
    offset = 1
    order = "string"
    severity = "string"
    sort_by = "string"
}

output "catalystcenter_network_bugs_results_bugs_example" {
    value = data.catalystcenter_network_bugs_results_bugs.example.items
}
