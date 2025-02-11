
data "catalystcenter_device_enrichment_details" "example" {
    provider = catalystcenter
    entity_type = "string"
    entity_value = "string"
}

output "catalystcenter_device_enrichment_details_example" {
    value = data.catalystcenter_device_enrichment_details.example.items
}
