
data "catalystcenter_system_issue_definitions_count" "example" {
  provider      = catalystcenter
  device_type   = "string"
  id            = "string"
  issue_enabled = "false"
  name          = "string"
  priority      = "string"
  profile_id    = "string"
  xca_lle_rid   = "string"
}

output "catalystcenter_system_issue_definitions_count_example" {
  value = data.catalystcenter_system_issue_definitions_count.example.item
}
