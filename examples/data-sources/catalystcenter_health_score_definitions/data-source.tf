
data "catalystcenter_health_score_definitions" "example" {
    provider = catalystcenter
    attribute = "string"
    device_type = "string"
    id = "string"
    include_for_overall_health = "false"
    limit = 1
    offset = 1
    xca_lle_rid = "string"
}

output "catalystcenter_health_score_definitions_example" {
    value = data.catalystcenter_health_score_definitions.example.items
}
