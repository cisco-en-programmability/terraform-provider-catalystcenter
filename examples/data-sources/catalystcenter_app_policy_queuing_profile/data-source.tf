
data "catalystcenter_app_policy_queuing_profile" "example" {
  provider = catalystcenter
  name     = "string"
}

output "catalystcenter_app_policy_queuing_profile_example" {
  value = data.catalystcenter_app_policy_queuing_profile.example.items
}
