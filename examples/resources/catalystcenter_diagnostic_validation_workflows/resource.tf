
resource "catalystcenter_diagnostic_validation_workflows" "example" {
    provider = catalystcenter

    parameters {

      description = "string"
      id = "string"
      name = "string"
      validation_set_ids = ["string"]
    }
}

output "catalystcenter_diagnostic_validation_workflows_example" {
    value = catalystcenter_diagnostic_validation_workflows.example
}
