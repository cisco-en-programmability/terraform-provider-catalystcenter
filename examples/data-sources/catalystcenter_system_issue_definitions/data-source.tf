
data "catalystcenter_system_issue_definitions" "example" {
    provider = catalystcenter
    attribute = "string"
    device_type = "string"
    id = "string"
    issue_enabled = "false"
    limit = 1
    name = "string"
    offset = 1
    order = "string"
    priority = "string"
    profile_id = "string"
    sort_by = "string"
    xca_lle_rid = "string"
}

output "catalystcenter_system_issue_definitions_example" {
    value = data.catalystcenter_system_issue_definitions.example.items
}

data "catalystcenter_system_issue_definitions" "example" {
    provider = catalystcenter
    id = "string"
    xca_lle_rid = "string"
}

output "catalystcenter_system_issue_definitions_example" {
    value = data.catalystcenter_system_issue_definitions.example.item
}
