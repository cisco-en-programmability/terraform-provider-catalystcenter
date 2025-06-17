
resource "catalystcenter_health_score_definitions" "example" {
  provider = catalystcenter
 
  parameters {

    id                             = "string"
    include_for_overall_health     = "false"
    synchronize_to_issue_threshold = "false"
    threshold_value                = 1.0
  }
}

output "catalystcenter_health_score_definitions_example" {
  value = catalystcenter_health_score_definitions.example
}