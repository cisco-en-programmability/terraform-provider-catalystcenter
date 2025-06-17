
resource "catalystcenter_system_issue_definitions_count" "example" {
  provider = catalystcenter
}

output "catalystcenter_system_issue_definitions_count_example" {
  value = catalystcenter_system_issue_definitions_count.example
}