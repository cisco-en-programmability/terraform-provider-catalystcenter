
data "catalystcenter_flexible_report_schedule" "example" {
  provider  = catalystcenter
  report_id = "string"
}

output "catalystcenter_flexible_report_schedule_example" {
  value = data.catalystcenter_flexible_report_schedule.example.item
}
