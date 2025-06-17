
data "catalystcenter_analytics_anc_policies" "example" {
  provider = catalystcenter
}

output "catalystcenter_analytics_anc_policies_example" {
  value = data.catalystcenter_analytics_anc_policies.example.items
}
