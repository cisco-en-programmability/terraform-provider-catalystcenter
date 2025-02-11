
data "catalystcenter_transit_network_health_summaries_count" "example" {
    provider = catalystcenter
    end_time = 1609459200
    id = "string"
    start_time = 1609459200
    xca_lle_rid = "string"
}

output "catalystcenter_transit_network_health_summaries_count_example" {
    value = data.catalystcenter_transit_network_health_summaries_count.example.items
}
