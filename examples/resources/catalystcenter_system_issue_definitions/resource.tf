
resource "catalystcenter_system_issue_definitions" "example" {
  provider = catalystcenter
 
  parameters {

    id                              = "string"
    issue_enabled                   = "false"
    priority                        = "string"
    synchronize_to_health_threshold = "false"
    threshold_value                 = 1.0
  }
}

output "catalystcenter_system_issue_definitions_example" {
  value = catalystcenter_system_issue_definitions.example
}