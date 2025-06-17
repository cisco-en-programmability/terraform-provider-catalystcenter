
data "catalystcenter_fabric_site_health_summaries_id" "example" {
  provider    = catalystcenter
  attribute   = "string"
  end_time    = 1609459200
  id          = "string"
  start_time  = 1609459200
  view        = "string"
  xca_lle_rid = "string"
}

output "catalystcenter_fabric_site_health_summaries_id_example" {
  value = data.catalystcenter_fabric_site_health_summaries_id.example.item
}
