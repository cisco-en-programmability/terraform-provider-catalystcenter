
data "catalystcenter_app_policy_default" "example" {
    provider = catalystcenter
}

output "catalystcenter_app_policy_default_example" {
    value = data.catalystcenter_app_policy_default.example.items
}
