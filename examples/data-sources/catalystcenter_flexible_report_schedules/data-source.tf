
data "catalystcenter_flexible_report_schedules" "example" {
  provider = catalystcenter
}

output "catalystcenter_flexible_report_schedules_example" {
  value = data.catalystcenter_flexible_report_schedules.example.items
}
