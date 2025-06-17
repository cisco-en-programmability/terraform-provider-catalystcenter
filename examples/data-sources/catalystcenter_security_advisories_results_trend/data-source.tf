
data "catalystcenter_security_advisories_results_trend" "example" {
  provider  = catalystcenter
  limit     = 1
  offset    = 1
  scan_time = 1.0
}

output "catalystcenter_security_advisories_results_trend_example" {
  value = data.catalystcenter_security_advisories_results_trend.example.items
}
