
data "catalystcenter_virtual_network_health_summaries_id_trend_analytics" "example" {
  provider       = catalystcenter
  attribute      = "string"
  end_time       = 1609459200
  id             = "string"
  limit          = 1
  offset         = 1
  order          = "string"
  start_time     = 1609459200
  trend_interval = "string"
  xca_lle_rid    = "string"
}

output "catalystcenter_virtual_network_health_summaries_id_trend_analytics_example" {
  value = data.catalystcenter_virtual_network_health_summaries_id_trend_analytics.example.items
}
