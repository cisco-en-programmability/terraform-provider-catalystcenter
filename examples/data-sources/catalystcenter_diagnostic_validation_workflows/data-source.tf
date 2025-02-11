
data "catalystcenter_diagnostic_validation_workflows" "example" {
    provider = catalystcenter
    end_time = 1609459200
    limit = 1
    offset = 1
    run_status = "string"
    start_time = 1609459200
}

output "catalystcenter_diagnostic_validation_workflows_example" {
    value = data.catalystcenter_diagnostic_validation_workflows.example.items
}

data "catalystcenter_diagnostic_validation_workflows" "example" {
    provider = catalystcenter
    id = "string"
}

output "catalystcenter_diagnostic_validation_workflows_example" {
    value = data.catalystcenter_diagnostic_validation_workflows.example.item
}
