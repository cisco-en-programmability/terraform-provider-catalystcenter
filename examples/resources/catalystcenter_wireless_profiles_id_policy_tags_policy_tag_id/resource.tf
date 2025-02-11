
resource "catalystcenter_wireless_profiles_id_policy_tags_policy_tag_id" "example" {
    provider = catalystcenter

    parameters {

      ap_zones = ["string"]
      id = "string"
      policy_tag_id = "string"
      policy_tag_name = "string"
      site_ids = ["string"]
    }
}

output "catalystcenter_wireless_profiles_id_policy_tags_policy_tag_id_example" {
    value = catalystcenter_wireless_profiles_id_policy_tags_policy_tag_id.example
}
