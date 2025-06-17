
data "catalystcenter_security_advisories_results_trend_count" "example" {
  provider  = catalystcenter
  scan_time = 1.0
}

output "catalystcenter_security_advisories_results_trend_count_example" {
  value = data.catalystcenter_security_advisories_results_trend_count.example.item
}
