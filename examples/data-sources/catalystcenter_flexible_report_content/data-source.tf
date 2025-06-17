
data "catalystcenter_flexible_report_content" "example" {
  provider     = catalystcenter
  execution_id = "string"
  report_id    = "string"
}

output "catalystcenter_flexible_report_content_example" {
  value = data.catalystcenter_flexible_report_content.example.item
}
