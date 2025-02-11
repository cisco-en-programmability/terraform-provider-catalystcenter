
data "catalystcenter_field_notices_results_trend_count" "example" {
    provider = catalystcenter
    scan_time = 1.0
}

output "catalystcenter_field_notices_results_trend_count_example" {
    value = data.catalystcenter_field_notices_results_trend_count.example.item
}
