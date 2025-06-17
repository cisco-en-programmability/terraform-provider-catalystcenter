
data "catalystcenter_field_notices_results_trend" "example" {
  provider  = catalystcenter
  limit     = 1
  offset    = 1
  scan_time = 1.0
}

output "catalystcenter_field_notices_results_trend_example" {
  value = data.catalystcenter_field_notices_results_trend.example.items
}
