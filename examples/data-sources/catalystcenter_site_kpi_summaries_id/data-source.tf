
data "catalystcenter_site_kpi_summaries_id" "example" {
  provider         = catalystcenter
  attribute        = "string"
  band             = "string"
  end_time         = 1609459200
  failure_category = "string"
  failure_reason   = "string"
  id               = "string"
  ssid             = "string"
  start_time       = 1609459200
  task_id          = "string"
  view             = "string"
  xca_lle_rid      = "string"
}

output "catalystcenter_site_kpi_summaries_id_example" {
  value = data.catalystcenter_site_kpi_summaries_id.example.item
}
