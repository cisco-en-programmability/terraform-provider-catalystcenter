
data "catalystcenter_intent_custom_issue_definitions" "example" {
  provider    = catalystcenter
  id          = "string"
  xca_lle_rid = "string"
}

output "catalystcenter_intent_custom_issue_definitions_example" {
  value = data.catalystcenter_intent_custom_issue_definitions.example.item
}
