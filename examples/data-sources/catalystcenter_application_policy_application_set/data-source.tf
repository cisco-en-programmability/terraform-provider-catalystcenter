
data "catalystcenter_application_policy_application_set" "example" {
  provider   = catalystcenter
  attributes = "string"
  limit      = 1
  name       = "string"
  offset     = 1
}

output "catalystcenter_application_policy_application_set_example" {
  value = data.catalystcenter_application_policy_application_set.example.items
}
