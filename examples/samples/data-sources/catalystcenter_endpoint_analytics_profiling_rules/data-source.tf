terraform {
  required_providers {
    catalystcenter = {
      version = "1.0.0-beta"
      source  = "hashicorp.com/edu/catalystcenter"
      # "hashicorp.com/edu/catalystcenter" is the local built source, change to "cisco-en-programmability/catalystcenter" to use downloaded version from registry
    }
  }
}

provider "catalystcenter" {
}
/*
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

*/
data "catalystcenter_endpoint_analytics_profiling_rules" "example" {
  provider = catalystcenter
  rule_id  = "string"
}

output "catalystcenter_endpoint_analytics_profiling_rules_example" {
  value = data.catalystcenter_endpoint_analytics_profiling_rules.example.item
}
