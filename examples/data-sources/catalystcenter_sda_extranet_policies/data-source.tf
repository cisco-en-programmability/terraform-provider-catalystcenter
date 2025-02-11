
data "catalystcenter_sda_extranet_policies" "example" {
    provider = catalystcenter
    extranet_policy_name = "string"
    limit = 1
    offset = 1
}

output "catalystcenter_sda_extranet_policies_example" {
    value = data.catalystcenter_sda_extranet_policies.example.items
}
