
data "catalystcenter_wireless_profiles_id_site_tags_site_tag_id" "example" {
  provider    = catalystcenter
  id          = "string"
  site_tag_id = "string"
}

output "catalystcenter_wireless_profiles_id_site_tags_site_tag_id_example" {
  value = data.catalystcenter_wireless_profiles_id_site_tags_site_tag_id.example.item
}
