
data "catalystcenter_custom_issue_definitions" "example" {
    provider = catalystcenter
    facility = "string"
    id = "string"
    is_enabled = "false"
    limit = 1
    mnemonic = "string"
    name = "string"
    offset = 1
    order = "string"
    priority = "string"
    profile_id = "string"
    severity = 1.0
    sort_by = "string"
}

output "catalystcenter_custom_issue_definitions_example" {
    value = data.catalystcenter_custom_issue_definitions.example.items
}

data "catalystcenter_custom_issue_definitions" "example" {
    provider = catalystcenter
    id = "string"
    xca_lle_rid = "string"
}

output "catalystcenter_custom_issue_definitions_example" {
    value = data.catalystcenter_custom_issue_definitions.example.item
}
