
data "catalystcenter_fabric_site_health_summaries_count" "example" {
  provider          = catalystcenter
  end_time          = 1609459200
  id                = "string"
  site_hierarchy    = "string"
  site_hierarchy_id = "string"
  start_time        = 1609459200
  xca_lle_rid       = "string"
}

output "catalystcenter_fabric_site_health_summaries_count_example" {
  value = data.catalystcenter_fabric_site_health_summaries_count.example.item
}
