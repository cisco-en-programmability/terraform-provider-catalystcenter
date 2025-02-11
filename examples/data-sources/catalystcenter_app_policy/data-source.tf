
data "catalystcenter_app_policy" "example" {
    provider = catalystcenter
    policy_scope = "string"
}

output "catalystcenter_app_policy_example" {
    value = data.catalystcenter_app_policy.example.items
}
