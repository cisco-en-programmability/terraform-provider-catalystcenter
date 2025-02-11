
data "catalystcenter_discovery_range" "example" {
    provider = catalystcenter
    records_to_return = 1
    start_index = 1
}

output "catalystcenter_discovery_range_example" {
    value = data.catalystcenter_discovery_range.example.items
}
