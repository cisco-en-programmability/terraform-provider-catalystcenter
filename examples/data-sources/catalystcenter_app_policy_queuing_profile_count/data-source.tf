
data "catalystcenter_app_policy_queuing_profile_count" "example" {
    provider = catalystcenter
}

output "catalystcenter_app_policy_queuing_profile_count_example" {
    value = data.catalystcenter_app_policy_queuing_profile_count.example.item
}
