
resource "catalystcenter_flexible_report_schedule" "example" {
  provider = catalystcenter
  parameters {

    report_id = "string"
    schedule  = "------"
  }
}

output "catalystcenter_flexible_report_schedule_example" {
  value = catalystcenter_flexible_report_schedule.example
}