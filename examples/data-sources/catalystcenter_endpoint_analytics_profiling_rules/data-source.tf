
data "catalystcenter_endpoint_analytics_profiling_rules" "example" {
  provider        = catalystcenter
  include_deleted = "false"
  limit           = 1
  offset          = 1
  order           = "string"
  rule_type       = "string"
  sort_by         = "string"
}

output "catalystcenter_endpoint_analytics_profiling_rules_example" {
  value = data.catalystcenter_endpoint_analytics_profiling_rules.example.items
}

data "catalystcenter_endpoint_analytics_profiling_rules" "example" {
  provider = catalystcenter
  rule_id  = "string"
}

output "catalystcenter_endpoint_analytics_profiling_rules_example" {
  value = data.catalystcenter_endpoint_analytics_profiling_rules.example.item
}
