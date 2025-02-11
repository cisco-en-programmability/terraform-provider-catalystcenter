
data "catalystcenter_custom_issue_definitions_count" "example" {
    provider = catalystcenter
    facility = "string"
    id = "string"
    is_enabled = "false"
    mnemonic = "string"
    name = "string"
    priority = "string"
    profile_id = "string"
    severity = 1.0
    xca_lle_rid = "string"
}

output "catalystcenter_custom_issue_definitions_count_example" {
    value = data.catalystcenter_custom_issue_definitions_count.example.item
}
