
data "catalystcenter_site_kpi_summaries_summary_analytics" "example" {
  provider    = catalystcenter
  task_id     = "string"
  xca_lle_rid = "string"
}

output "catalystcenter_site_kpi_summaries_summary_analytics_example" {
  value = data.catalystcenter_site_kpi_summaries_summary_analytics.example.item
}
