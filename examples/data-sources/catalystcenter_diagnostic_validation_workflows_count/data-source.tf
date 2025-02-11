
data "catalystcenter_diagnostic_validation_workflows_count" "example" {
    provider = catalystcenter
    end_time = 1609459200
    run_status = "string"
    start_time = 1609459200
}

output "catalystcenter_diagnostic_validation_workflows_count_example" {
    value = data.catalystcenter_diagnostic_validation_workflows_count.example.item
}
