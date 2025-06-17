
data "catalystcenter_transit_network_health_summaries" "example" {
  provider    = catalystcenter
  attribute   = "string"
  end_time    = 1609459200
  id          = "string"
  limit       = 1
  offset      = 1
  order       = "string"
  sort_by     = "string"
  start_time  = 1609459200
  view        = "string"
  xca_lle_rid = "string"
}

output "catalystcenter_transit_network_health_summaries_example" {
  value = data.catalystcenter_transit_network_health_summaries.example.items
}
