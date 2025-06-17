
resource "catalystcenter_assurance_issues_resolve" "example" {
  provider    = catalystcenter
  xca_lle_rid = "string"
  parameters = [{

    issue_ids = ["string"]
  }]
}

output "catalystcenter_assurance_issues_resolve_example" {
  value = catalystcenter_assurance_issues_resolve.example
}