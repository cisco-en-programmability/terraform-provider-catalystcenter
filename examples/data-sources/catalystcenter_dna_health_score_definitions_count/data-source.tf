
data "catalystcenter_dna_health_score_definitions_count" "example" {
  provider                   = catalystcenter
  device_type                = "string"
  id                         = "string"
  include_for_overall_health = "false"
  xca_lle_rid                = "string"
}

output "catalystcenter_dna_health_score_definitions_count_example" {
  value = data.catalystcenter_dna_health_score_definitions_count.example.item
}
