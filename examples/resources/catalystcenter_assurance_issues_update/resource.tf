
resource "catalystcenter_assurance_issues_update" "example" {
  provider        = catalystcenter
  accept_language = "string"
  id              = "string"
  xca_lle_rid     = "string"
  parameters = [{

    notes = "string"
  }]
}

output "catalystcenter_assurance_issues_update_example" {
  value = catalystcenter_assurance_issues_update.example
}