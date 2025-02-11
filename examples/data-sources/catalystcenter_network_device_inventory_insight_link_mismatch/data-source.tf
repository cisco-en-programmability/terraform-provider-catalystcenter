
data "catalystcenter_network_device_inventory_insight_link_mismatch" "example" {
    provider = catalystcenter
    category = "string"
    limit = 1
    offset = 1
    order = "string"
    site_id = "string"
    sort_by = "string"
}

output "catalystcenter_network_device_inventory_insight_link_mismatch_example" {
    value = data.catalystcenter_network_device_inventory_insight_link_mismatch.example.items
}
