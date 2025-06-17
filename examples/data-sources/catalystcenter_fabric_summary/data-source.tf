
data "catalystcenter_fabric_summary" "example" {
  provider          = catalystcenter
  end_time          = 1609459200
  site_hierarchy    = "string"
  site_hierarchy_id = "string"
  start_time        = 1609459200
  xca_lle_rid       = "string"
}

output "catalystcenter_fabric_summary_example" {
  value = data.catalystcenter_fabric_summary.example.item
}
