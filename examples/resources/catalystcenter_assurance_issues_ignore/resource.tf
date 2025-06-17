
resource "catalystcenter_assurance_issues_ignore" "example" {
  provider    = catalystcenter
  xca_lle_rid = "string"
  parameters = [{

    ignore_hours = 1
    issue_ids    = ["string"]
  }]
}

output "catalystcenter_assurance_issues_ignore_example" {
  value = catalystcenter_assurance_issues_ignore.example
}