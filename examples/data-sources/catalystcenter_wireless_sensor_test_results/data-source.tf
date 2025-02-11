
data "catalystcenter_wireless_sensor_test_results" "example" {
    provider = catalystcenter
    end_time = 1609459200
    site_id = "string"
    start_time = 1609459200
    test_failure_by = "string"
}

output "catalystcenter_wireless_sensor_test_results_example" {
    value = data.catalystcenter_wireless_sensor_test_results.example.item
}
