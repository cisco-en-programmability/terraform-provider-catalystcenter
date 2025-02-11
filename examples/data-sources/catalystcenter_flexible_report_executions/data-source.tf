
data "catalystcenter_flexible_report_executions" "example" {
    provider = catalystcenter
    report_id = "string"
}

output "catalystcenter_flexible_report_executions_example" {
    value = data.catalystcenter_flexible_report_executions.example.item
}
