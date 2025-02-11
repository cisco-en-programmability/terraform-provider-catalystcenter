
resource "catalystcenter_wireless_profiles_id_site_tags_site_tag_id" "example" {
    provider = catalystcenter

    parameters {

      ap_profile_name = "string"
      flex_profile_name = "string"
      id = "string"
      site_ids = ["string"]
      site_tag_id = "string"
      site_tag_name = "string"
    }
}

output "catalystcenter_wireless_profiles_id_site_tags_site_tag_id_example" {
    value = catalystcenter_wireless_profiles_id_site_tags_site_tag_id.example
}
