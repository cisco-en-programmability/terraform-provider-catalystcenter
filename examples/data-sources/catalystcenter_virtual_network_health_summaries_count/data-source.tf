
data "catalystcenter_virtual_network_health_summaries_count" "example" {
    provider = catalystcenter
    end_time = 1609459200
    id = "string"
    start_time = 1609459200
    vn_layer = "string"
    xca_lle_rid = "string"
}

output "catalystcenter_virtual_network_health_summaries_count_example" {
    value = data.catalystcenter_virtual_network_health_summaries_count.example.item
}
