
data "catalystcenter_reports_executions" "example" {
  provider  = catalystcenter
  report_id = "string"
}

output "catalystcenter_reports_executions_example" {
  value = data.catalystcenter_reports_executions.example.item
}
