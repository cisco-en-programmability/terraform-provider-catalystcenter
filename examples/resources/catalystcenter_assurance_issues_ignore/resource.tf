
resource "catalystcenter_assurance_issues_ignore" "example" {
  provider    = meraki
  xca_lle_rid = "string"
  parameters {

    issue_ids = ["string"]
  }
}

output "catalystcenter_assurance_issues_ignore_example" {
  value = catalystcenter_assurance_issues_ignore.example
}