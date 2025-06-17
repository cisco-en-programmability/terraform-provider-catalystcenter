
data "catalystcenter_issues_enrichment_details" "example" {
  provider     = catalystcenter
  entity_type  = "string"
  entity_value = "string"
}

output "catalystcenter_issues_enrichment_details_example" {
  value = data.catalystcenter_issues_enrichment_details.example.item
}
