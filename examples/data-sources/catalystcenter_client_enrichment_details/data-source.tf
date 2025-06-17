
data "catalystcenter_client_enrichment_details" "example" {
  provider       = catalystcenter
  entity_type    = "string"
  entity_value   = "string"
  issue_category = "string"
}

output "catalystcenter_client_enrichment_details_example" {
  value = data.catalystcenter_client_enrichment_details.example.items
}
