
data "catalystcenter_application_policy_application_set_count" "example" {
  provider            = catalystcenter
  scalable_group_type = "string"
}

output "catalystcenter_application_policy_application_set_count_example" {
  value = data.catalystcenter_application_policy_application_set_count.example.item
}
