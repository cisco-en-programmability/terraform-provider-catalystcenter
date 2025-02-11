
data "catalystcenter_wireless_profiles_id_policy_tags_policy_tag_id" "example" {
    provider = catalystcenter
    id = "string"
    policy_tag_id = "string"
}

output "catalystcenter_wireless_profiles_id_policy_tags_policy_tag_id_example" {
    value = data.catalystcenter_wireless_profiles_id_policy_tags_policy_tag_id.example.item
}
