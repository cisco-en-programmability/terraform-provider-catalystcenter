
resource "catalystcenter_flexible_report_execute" "example" {
  provider  = meraki
  report_id = "string"
  parameters {

  }
}

output "catalystcenter_flexible_report_execute_example" {
  value = catalystcenter_flexible_report_execute.example
}