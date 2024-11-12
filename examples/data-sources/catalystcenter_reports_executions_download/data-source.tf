
data "catalystcenter_reports_executions_download" "example" {
  provider     = catalystcenter
  execution_id = "string"
  report_id    = "string"
}

output "catalystcenter_reports_executions_download_example" {
  value = data.catalystcenter_reports_executions_download.example.item
}
