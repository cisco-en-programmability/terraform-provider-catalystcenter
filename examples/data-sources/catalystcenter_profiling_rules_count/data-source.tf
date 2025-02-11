
data "catalystcenter_profiling_rules_count" "example" {
    provider = catalystcenter
    include_deleted = "false"
    rule_type = "string"
}

output "catalystcenter_profiling_rules_count_example" {
    value = data.catalystcenter_profiling_rules_count.example.item
}
